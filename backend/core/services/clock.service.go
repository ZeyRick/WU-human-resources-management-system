package services

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/clock"
	clocksetting "backend/core/models/clock_setting"
	"backend/core/models/employee"
	"backend/core/types"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/variable"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ClockService struct {
	repo     clock.ClockRepo
	emp      *employee.EmployeeRepo
	clockset clocksetting.ClockSettingRepo
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

func (srv *ClockService) ClockLocation(longtitude float64, latitude float64) (bool, error) {
	clockLocation, err := srv.clockset.Get()
	if err != nil {
		return false, err
	}
	clockcoordinate := strings.SplitN(clockLocation.Coordinate, ",", 2)
	xCoordinate, err := strconv.ParseFloat(clockcoordinate[0], 64)
	if err != nil {
		return false, err
	}
	yCoordinate, err := strconv.ParseFloat(clockcoordinate[1], 64)
	if err != nil {
		return false, err
	}
	distance := math.Sqrt(math.Pow(latitude-xCoordinate, 2) + math.Pow(longtitude-yCoordinate, 2))
	if distance > float64(*clockLocation.ClockRange) {
		return false, nil
	}
	return true, nil
}
