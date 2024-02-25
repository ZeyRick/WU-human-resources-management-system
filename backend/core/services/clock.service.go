package services

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/clock"
	"backend/core/models/employee"
	"backend/core/types"
	"backend/pkg/excelhelper"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/logger"
	"backend/pkg/times"
	"backend/pkg/variable"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

type ClockService struct {
	repo clock.ClockRepo
	emp  *employee.EmployeeRepo
}

func NewClockService() *ClockService {
	return &ClockService{
		repo: *clock.NewClockRepo(),
	}
}

func (srv *ClockService) Clock(w http.ResponseWriter, r *http.Request, payload dtos.Clock) error {
	if payload.ClockType == types.ClockOut {
		prevClock, err := srv.repo.LatestClockIn(payload.EmployeeId)
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				https.ResponseError(w, r, http.StatusInternalServerError, "You must clock in first before clock out")
				return err
			}
			helper.UnexpectedError(w, r, err)
			return err
		}
		curTime := time.Now().UTC()
		minuteWork := int(math.Round(prevClock.CreatedAt.Sub(curTime).Minutes()))
		err = srv.repo.Create(
			&clock.Clock{
				EmployeeId:     payload.EmployeeId,
				ClockType:      payload.ClockType,
				BaseModel:      models.BaseModel{CreatedAt: curTime},
				ClockOutMinute: &minuteWork,
				ClockInId:      variable.Create[int](int(prevClock.ID))})
		if err != nil {
			helper.UnexpectedError(w, r, err)
			return err
		}
		return nil
	}
	err := srv.repo.Create(&clock.Clock{EmployeeId: payload.EmployeeId, ClockType: payload.ClockType})
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return err
	}
	return nil
}

func (srv *ClockService) List(pageOpt *dtos.PageOpt, dto *dtos.ClockFilter) (*types.ListData[clock.Clock], error) {
	result, err := srv.repo.List(pageOpt, dto)
	return result, err
}

func (srv *ClockService) Attendence(pageOpt *dtos.PageOpt, dto *dtos.AttendenceFilter) (*types.ListData[clock.Clock], error) {
	result, err := srv.repo.Attendence(pageOpt, dto)
	return result, err
}

func (srv *ClockService) AttendenceExport(w http.ResponseWriter, r *http.Request, pageOpt *dtos.PageOpt, dto *dtos.AttendenceFilter) {
	result, err := srv.repo.Attendence(pageOpt, dto)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	columns := []string{
		"Date",
		"Employee Name",
		"Clock In Time",
		"Clock Out Time",
		"Total Work Minute",
		"Work Time",
		"Status",
	}
	sheetName := "Attendence"
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName)
	defer func() {
		if err := f.Close(); err != nil {
			logger.Trace(err)
		}
	}()

	for index, col := range columns {
		colIndex := variable.IntToAlphabet(index)
		f.SetCellValue(sheetName, fmt.Sprintf("%s1", colIndex), col)
	}

	for index, attendence := range *result.Data {
		rowIndex := index + 2
		colIndex := 0
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, attendence.CreatedAt.Format("2006-01-02"))
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, attendence.Employee.Name)
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, attendence.ClockIn.CreatedAt.Format("15:04:05"))
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, attendence.CreatedAt.Format("15:04:05"))
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, *attendence.ClockOutMinute)
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, fmt.Sprintf("%s-%s", attendence.Schedule.ClockInTime.Format("15:04:05"), attendence.Schedule.ClockOutTime.Format("15:04:05")))

		logger.Console(attendence.Employee.Name)
		isLate, err := times.IsTimeAfter(attendence.ClockIn.CreatedAt, attendence.Schedule.ClockInTime)
		if err != nil {
			helper.UnexpectedError(w, r, err)
			return
		}
		isEarly, err := times.IsTimeBefore(attendence.CreatedAt, attendence.Schedule.ClockOutTime)
		if err != nil {
			helper.UnexpectedError(w, r, err)
			return
		}
		var status string
		var color string
		if isLate && isEarly {
			status = "Late-Early"
			color = "#730000"
		} else if isLate {
			status = "Late"
			color = "#ff0000"
		} else if isEarly {
			status = "Early"
			color = "#0022ba"
		} else {
			status = "On Time"
			color = "#00bd00"
		}
		style, err := f.NewStyle(&excelize.Style{
			Font: &excelize.Font{
				Color: color,
			},
		})
		if err != nil {
			helper.UnexpectedError(w, r, err)
			return
		}
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, status, style)
	}

	var startDate string = ""
	var endDate string = ""
	if dto.StartDate != "" && dto.EndDate != "" {
		startDate = strings.ReplaceAll(dto.StartDate, "-", "_")
		endDate = strings.ReplaceAll(dto.EndDate, "-", "_")
		startDate = strings.ReplaceAll(startDate, " ", "_")
		endDate = strings.ReplaceAll(endDate, " ", "_")
	}
	fileName := fmt.Sprintf("Attendence_%s_%s.xlsx", startDate, endDate)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	f.Write(w)
}

func (srv *ClockService) ClockFromTelegram(telegramID *int64, clockType types.ClockType) error {
	employee, err := srv.emp.FindTelegramId(telegramID)
	empID := int(employee.ID)
	empID2 := &empID
	if err != nil {
		return err
	}
	if clockType == types.ClockOut {
		prevClock, err := srv.repo.LatestClockIn(empID2)
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				return err
			}
			return err
		}
		curTime := time.Now().UTC()
		hourWork := int(math.Round(prevClock.CreatedAt.Sub(curTime).Hours()))
		err = srv.repo.Create(&clock.Clock{EmployeeId: empID2, ClockType: clockType, BaseModel: models.BaseModel{CreatedAt: curTime}, ClockOutMinute: &hourWork})
		if err != nil {
			return err
		}
		return nil
	}
	err = srv.repo.Create(&clock.Clock{EmployeeId: empID2, ClockType: clockType})
	if err != nil {
		return err
	}
	return nil
}
