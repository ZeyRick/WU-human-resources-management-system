package services

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/clock"
	clocksetting "backend/core/models/clock_setting"
	"backend/core/models/employee"
	"backend/core/models/schedule"
	"backend/core/types"
	"backend/pkg/excelhelper"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/logger"
	"backend/pkg/variable"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

type ClockService struct {
	repo         *clock.ClockRepo
	emp          *employee.EmployeeRepo
	clockset     *clocksetting.ClockSettingRepo
	scheduleRepo *schedule.ScheduleRepo
}

func NewClockService() *ClockService {
	return &ClockService{
		repo:         clock.NewClockRepo(),
		emp:          employee.NewEmployeeRepo(),
		clockset:     clocksetting.NewClockSettingRepo(),
		scheduleRepo: schedule.NewScheduleRepo(),
	}
}

func (srv *ClockService) Clock(w http.ResponseWriter, r *http.Request, payload dtos.Clock) error {
	if payload.ClockType == types.ClockOut {
		prevClock, err := srv.repo.LatestClock(payload.EmployeeId)
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				https.ResponseError(w, r, http.StatusInternalServerError, "You must clock in first before clock out")
				return err
			}
			helper.UnexpectedError(w, r, err)
			return err
		}
		curTime := time.Now().UTC()
		minuteWork := int(math.Round(math.Abs(curTime.Sub(prevClock.CreatedAt).Minutes())))
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

		isEarly := attendence.Status != ""
		isLate := attendence.ClockIn.Status != ""
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
		prevClock, err := srv.repo.LatestClock(empID2)
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

func (srv *ClockService) CheckAvaiableClockIn(employeeId *int) (string, error) {
	// Check if today has already clock in
	prevClock, err := srv.repo.LatestClockIn(employeeId)
	if err != nil && !strings.Contains(err.Error(), "record not found") {
		return "", err
	}
	if prevClock != nil {
		prevClockYear, prevClockMonth, prevClockDay := prevClock.CreatedAt.Add(7 * time.Hour).Date()
		curYear, curMonth, curDay := time.Now().UTC().Add(7 * time.Hour).Date()
		if prevClockDay == curDay && prevClockYear == curYear && prevClockMonth == curMonth {
			return "You have already clocked in today.", nil
		}

	}
	return "", nil
}

func (srv *ClockService) CheckAvaiableClockOut(employeeId *int) (string, error) {
	// check previous clock in
	prevClock, err := srv.repo.LatestClock(employeeId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return "You must clock in first before clock out", nil
		}
		return "", err
	}
	if prevClock.ClockType == types.ClockOut {
		return "You must clock in first before clock out", nil
	}
	return "", err
}

func (srv *ClockService) ClockIn(employeeId *int, longtitude float64, latitude float64) (string, error) {
	// Checking location
	clockSetting, err := srv.clockset.Get()
	if err != nil {
		return "", err
	}
	clockcoordinate := strings.SplitN(clockSetting.Coordinate, ",", 2)
	xCoordinate, err := strconv.ParseFloat(clockcoordinate[0], 64)
	if err != nil {
		return "", err
	}
	yCoordinate, err := strconv.ParseFloat(clockcoordinate[1], 64)
	if err != nil {
		return "", err
	}
	distance := math.Sqrt(math.Pow(latitude-xCoordinate, 2) + math.Pow(longtitude-yCoordinate, 2))
	if distance > float64(*clockSetting.ClockRange) {
		return "You are not inside clock range", nil
	}

	// Checking allow time
	utcTime := time.Now().UTC()
	scope := fmt.Sprintf("%d-%02d", utcTime.Year(), utcTime.Month())
	schedule, err := srv.scheduleRepo.FindExistedScope(employeeId, scope)
	if err != nil {
		return "", err
	}
	if schedule.ID == 0 {
		return "No schedule for you today", nil
	}
	utcPlus7 := utcTime.Add(7 * time.Hour)
	compareTimeStr := fmt.Sprintf("%s-01 %02d:%02d:%02d", schedule.Scope, utcPlus7.Hour(), utcPlus7.Minute(), utcPlus7.Second())
	compareCurTime, err := time.Parse("2006-01-02 15:04:05", compareTimeStr)
	compareCurTime = compareCurTime.Add(-7 * time.Hour)
	if err != nil {
		return "", err
	}

	var status string = ""
	clockInTime := time.Now().UTC()
	differentMinutes := int(compareCurTime.Sub(schedule.ClockInTime).Minutes())
	if differentMinutes > 0 && differentMinutes <= *clockSetting.AllowTime {
		clockInTime = clockInTime.Add(time.Duration(differentMinutes * int(time.Minute)))
	} else if differentMinutes > *clockSetting.AllowTime {
		status = "late"
	}

	// Check if today has already clock in
	errMsg, err := srv.CheckAvaiableClockIn(employeeId)
	if err != nil {
		return "", err
	}
	if errMsg != "" {
		return errMsg, nil
	}

	err = srv.repo.Create(&clock.Clock{
		BaseModel: models.BaseModel{
			CreatedAt: clockInTime,
		},
		EmployeeId: employeeId,
		ClockType:  types.ClockIn,
		Status:     status,
		ScheduleId: variable.Create[int](int(schedule.ID))})
	if err != nil {
		return "", err
	}
	return "", nil
}

func (srv *ClockService) ClockOut(employeeId *int, longtitude float64, latitude float64) (string, error) {
	// Checking location
	clockSetting, err := srv.clockset.Get()
	if err != nil {
		return "", err
	}
	clockcoordinate := strings.SplitN(clockSetting.Coordinate, ",", 2)
	xCoordinate, err := strconv.ParseFloat(clockcoordinate[0], 64)
	if err != nil {
		return "", err
	}
	yCoordinate, err := strconv.ParseFloat(clockcoordinate[1], 64)
	if err != nil {
		return "", err
	}
	distance := math.Sqrt(math.Pow(latitude-xCoordinate, 2) + math.Pow(longtitude-yCoordinate, 2))
	if distance > float64(*clockSetting.ClockRange) {
		return "You are not inside clock range", nil
	}

	// Checking allow time
	utcTime := time.Now().UTC()
	scope := fmt.Sprintf("%d-%02d", utcTime.Year(), utcTime.Month())
	schedule, err := srv.scheduleRepo.FindExistedScope(employeeId, scope)
	if err != nil {
		return "", err
	}
	if schedule.ID == 0 {
		return "No schedule for you today", nil
	}
	utcPlus7 := utcTime.Add(7 * time.Hour)
	compareTimeStr := fmt.Sprintf("%s-01 %02d:%02d:%02d", schedule.Scope, utcPlus7.Hour(), utcPlus7.Minute(), utcPlus7.Second())
	compareCurTime, err := time.Parse("2006-01-02 15:04:05", compareTimeStr)
	compareCurTime = compareCurTime.Add(-7 * time.Hour)
	if err != nil {
		return "", err
	}

	curTime := time.Now().UTC()
	var status string = ""
	differentMinutes := int(schedule.ClockOutTime.Sub(compareCurTime).Minutes())
	if differentMinutes > 0 && differentMinutes <= *clockSetting.AllowTime {
		curTime = curTime.Add(time.Duration(differentMinutes * int(time.Minute)))
	} else if differentMinutes > *clockSetting.AllowTime {
		status = "early"
	}

	// check previous clock in
	prevClock, err := srv.repo.LatestClock(employeeId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return "You must clock in first before clock out", nil
		}
		return "", err
	}
	if prevClock.ClockType == types.ClockOut {
		return "You must clock in first before clock out", nil
	}
	minuteWork := int(math.Round(math.Abs(curTime.Sub(prevClock.CreatedAt).Minutes())))
	err = srv.repo.Create(
		&clock.Clock{
			EmployeeId:     employeeId,
			ClockType:      types.ClockOut,
			BaseModel:      models.BaseModel{CreatedAt: curTime},
			ClockOutMinute: &minuteWork,
			ClockInId:      variable.Create[int](int(prevClock.ID)),
			Status:         status,
			ScheduleId:     variable.Create[int](int(schedule.ID)),
		})
	if err != nil {
		return "", err
	}
	return "", nil
}
