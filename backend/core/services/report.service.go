package services

import (
	"backend/adapters/dtos"
	"backend/core/models/clock"
	"backend/core/types"
)

type ReportService struct {
	clockRepo clock.ClockRepo
}

func NewReportService() *ReportService {
	return &ReportService{
		clockRepo: *clock.NewClockRepo(),
	}
}

func (srv *ReportService) List(pageOpt *dtos.PageOpt, dto *dtos.ReportFilter) (*[]types.ClockReports, error) {
	return srv.clockRepo.SumReport(pageOpt, dto)
}
