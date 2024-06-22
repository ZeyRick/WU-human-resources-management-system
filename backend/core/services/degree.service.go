package services

import (
	"backend/adapters/dtos"
	"backend/core/models/degree"
	"backend/core/types"
	"backend/pkg/helper"
	"backend/pkg/https"
	"net/http"
	"strings"
)

type DegreeService struct {
	repo *degree.DegreeRepo
}

func NewDegreeService() *DegreeService {
	return &DegreeService{
		repo: degree.NewDegreeRepo(),
	}
}

func (srv *DegreeService) All() (*[]degree.Degree, error) {
	return srv.repo.All()
}

func (srv *DegreeService) List(pageOpt *dtos.PageOpt, dto *dtos.DegreeFilter) (*types.ListData[degree.Degree], error) {
	return srv.repo.List(pageOpt, dto)
}
func (srv *DegreeService) SearchList(dto *dtos.DegreeFilter) ([]*degree.Degree, error) {
	return srv.repo.SearchList(dto)
}
func (srv *DegreeService) Add(w http.ResponseWriter, r *http.Request, payload *dtos.AddDegree) {
	err := srv.repo.Create(&degree.Degree{
		Alias: payload.Alias,
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

func (srv *DegreeService) Edit(w http.ResponseWriter, r *http.Request, payload *dtos.AddDegree) {
	rowAffected, err := srv.repo.UpdateById(&degree.Degree{
		Alias: payload.Alias,
		ID:    payload.ID,
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
