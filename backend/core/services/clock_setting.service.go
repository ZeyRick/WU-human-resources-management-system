package services

import (
	"backend/core/models"
	"backend/core/repos"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/validate"
	"net/http"
)

type ClockSettingService struct {
	repo repos.ClockSettingRepo
}

func NewClockSettingService() *ClockSettingService {
	return &ClockSettingService{
		repo: *repos.NewClockSettingRepo(),
	}
}

func (svc *ClockSettingService) Get() (*models.ClockSetting, error) {
	return svc.repo.Get()
}

func (svc *ClockSettingService) Save(w http.ResponseWriter, r *http.Request, newClockSetting *models.ClockSetting) {
	if (!validate.IsValidGoogleCoordinate(newClockSetting.Coordinate)) {
		https.ResponseError(w,r, http.StatusBadRequest, "[" + newClockSetting.Coordinate + "] Is Not Valid Coordinate")
		return
	}
	effectedRow, err := svc.repo.Update(newClockSetting)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if (*effectedRow == 0) {
		err = svc.repo.Create(newClockSetting)
		if err != nil {
			helper.UnexpectedError(w, r, err)
			return
		}
	}
	https.ResponseMsg(w,r,http.StatusCreated, "Setting Saved")
}
