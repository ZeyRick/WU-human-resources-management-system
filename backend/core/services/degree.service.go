package services

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/repos"
	"backend/core/types"
	"backend/pkg/helper"
	"backend/pkg/https"
	"net/http"
	"strings"
)

type DegreeService struct {
	repo *repos.DegreeRepo
}

func NewDegreeService() *DegreeService {
	return &DegreeService{
		repo: repos.NewDegreeRepo(),
	}
}

func (srv *DegreeService) All() (*[]models.Degree, error) {
	return srv.repo.All()
}

func (srv *DegreeService) GetByEmployee(employeeId uint) ([]models.Degree, error) {
	return srv.repo.GetByEmployee(&employeeId)
}


func (srv *DegreeService) List(pageOpt *dtos.PageOpt, dto *dtos.DegreeFilter) (*types.ListData[models.Degree], error) {
	return srv.repo.List(pageOpt, dto)
}

func (srv *DegreeService) Add(w http.ResponseWriter, r *http.Request, payload *dtos.AddDegree) {
	err := srv.repo.Create(&models.Degree{
		Alias: payload.Alias,
		Rate:  payload.Rate,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			https.ResponseError(w, r, http.StatusBadRequest, "Course already existed")
			return
		}
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Degree created")
}

func (srv *DegreeService) Edit(w http.ResponseWriter, r *http.Request, payload *dtos.AddDegree, id uint) {
	rowAffected, err := srv.repo.UpdateById(&models.Degree{
		Alias: payload.Alias,
		ID:    id,
		Rate:  payload.Rate,
	})
	if err != nil || rowAffected <= 0 {
		if strings.Contains(err.Error(), "Duplicate entry") {
			https.ResponseError(w, r, http.StatusBadRequest, "Degree alias already existed")
			return
		}
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Degree updated")
}
