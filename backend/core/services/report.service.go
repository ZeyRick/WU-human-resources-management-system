package services

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/repos"
	"backend/core/types"
	"backend/pkg/excelhelper"
	"backend/pkg/helper"
	"backend/pkg/logger"
	"backend/pkg/variable"
	"fmt"
	"net/http"
	"strings"

	"github.com/xuri/excelize/v2"
)

type ReportService struct {
	clockRepo    repos.ClockRepo
	employeeRepo repos.EmployeeRepo
	degreeRepo   repos.DegreeRepo
	courseRepo   repos.CourseRepo
	userRepo     repos.UserRepo
}

func NewReportService() *ReportService {
	return &ReportService{
		clockRepo:    *repos.NewClockRepo(),
		employeeRepo: *repos.NewEmployeeRepo(),
		degreeRepo:   *repos.NewDegreeRepo(),
		courseRepo:   *repos.NewCourseRepo(),
	}
}

func (srv *ReportService) List(pageOpt *dtos.PageOpt, dto *dtos.ReportFilter) (*types.ListData[models.ClockReports], error) {
	return srv.clockRepo.SumReport(pageOpt, dto)
}

func (srv *ReportService) Export(w http.ResponseWriter, r *http.Request, pageOpt *dtos.PageOpt, dto *dtos.ReportFilter) {
	result, err := srv.clockRepo.SumReport(pageOpt, dto)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	columns := []string{
		"Employee ID",
		"Employee Name",
		"Course",
		"Total Work Minute",
		"Total Late Minute",
		"Total Early Minute",
	}
	sheetName := "Reports"
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

	for index, row := range *result.Data {
		rowIndex := index + 2
		colIndex := 0
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, *row.EmployeeId)
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, row.Name)
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, row.DegreeAlias)
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, *row.TotalWorkMinute)
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, *row.TotalLateMinute)
		excelhelper.SetCell(f, sheetName, &colIndex, rowIndex, *row.TotalEarlyMinute)
	}

	var startDate string = ""
	var endDate string = ""
	if dto.StartDate != "" && dto.EndDate != "" {
		startDate = strings.ReplaceAll(dto.StartDate, "-", "_")
		endDate = strings.ReplaceAll(dto.EndDate, "-", "_")
		startDate = strings.ReplaceAll(startDate, " ", "_")
		endDate = strings.ReplaceAll(endDate, " ", "_")
	}
	fileName := fmt.Sprintf("Reports_%s_%s.xlsx", startDate, endDate)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	f.Write(w)
}

func (srv *ReportService) Dashboard() (*types.DashboardSummary, error) {
	employeeCounts, err := srv.employeeRepo.CountByType()
	if err != nil {
		return nil, err
	}
	for _, emType := range types.EmployeeTypes {
		for _, emCount := range employeeCounts {
			if emCount.EmployeeType == emType {
				goto out
			}
		}
		employeeCounts = append(employeeCounts, types.EmployeeCountType{TotalCount: 0, EmployeeType: emType})
	out:
	}

	degreeCount, err := srv.degreeRepo.Count()
	if err != nil {
		return nil, err
	}

	courseCount, err := srv.courseRepo.Count()
	if err != nil {
		return nil, err
	}

	userCount, err := srv.userRepo.Count()
	if err != nil {
		return nil, err
	}

	result := types.DashboardSummary{
		EmployeeCounts: employeeCounts,
		DegreeCount:    degreeCount,
		CourseCount:    courseCount,
		UserCount:      userCount,
	}
	return &result, err
}
