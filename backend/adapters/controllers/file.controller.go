package controllers

import (
	"backend/pkg/helper"
	"backend/pkg/https"
	"net/http"
)

type FileController struct {
}

func NewFileController() *FileController {
	return &FileController{}
}

func (ctrl *FileController) GetImageFile(w http.ResponseWriter, r *http.Request) {
	fileName, err := https.GetParamsStr(r, "fileName")
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if fileName == "" {
		https.ResponseError(w, r, http.StatusBadRequest, "Missing file name")
		return
	}
	
}
