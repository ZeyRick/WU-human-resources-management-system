package excel

import (
	"backend/adapters/dtos"
	"backend/core/types"
	"bytes"
	"errors"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ReadEmployeeExcel(buffer *bytes.Buffer) ([]dtos.AddEmployee, error) {
	var employeesdot []dtos.AddEmployee
	file, err := excelize.OpenReader(buffer)
	if err != nil {
		return employeesdot, err
	}

	sheetName := "Sheet1"

	rows, err := file.GetRows(sheetName)

	if err != nil {
		return employeesdot, err
	}
	header := rows[0]
	targetCols := []string{"Name", "DepartmentId", "EmployeeType", "Salary"}
	var name []string
	var departmentId []int
	var employeeType []types.EmployeeType
	var salary []float64
	var checked bool
	for _, targetColName := range targetCols {
		checked = false
		for i, colName := range header {
			if colName == targetColName {
				switch targetColName {
				case "Name":
					for _, row := range rows[1:] {
						name = append(name, row[i])
					}
				case "DepartmentId":
					for _, row := range rows[1:] {
						departmentids, err := strconv.Atoi(row[i])
						if err != nil {
							return employeesdot, err
						}
						departmentId = append(departmentId, departmentids)
					}
				case "EmployeeType":
					for _, row := range rows[1:] {
						employeestype, err := stringToEmployeeType(row[i])
						if err != nil {
							return employeesdot, err
						}
						employeeType = append(employeeType, employeestype)
					}
				case "Salary":
					for _, row := range rows[1:] {
						salaries, err := strconv.ParseFloat(row[i], 64)
						if err != nil {
							return employeesdot, err
						}
						salary = append(salary, salaries)
					}
				}
				checked = true
			}
		}
		if !checked {
			if name == nil || departmentId == nil || employeeType == nil || salary == nil {
				return employeesdot, errors.New("empty data")
			}
		}
	}
	if len(name) == len(departmentId) && len(departmentId) == len(employeeType) && len(employeeType) == len(salary) {
		var employeedot dtos.AddEmployee
		for i := range name {
			employeedot.Name = name[i]
			employeedot.DepartmentId = departmentId[i]
			employeedot.EmployeeType = employeeType[i]
			employeedot.Salary = salary[i]
			employeesdot = append(employeesdot, employeedot)
		}
	}
	return employeesdot, nil
}

func stringToEmployeeType(employeeTypeString string) (types.EmployeeType, error) {
	switch employeeTypeString {
	case string(types.Fulltime):
		return types.Fulltime, nil
	case string(types.Parttime):
		return types.Parttime, nil
	default:
		return "", errors.New("invalid data type for conversion")
	}
}

func ReadUserExcel(buffer *bytes.Buffer) ([]dtos.UserRegister, error) {
	var usersdot []dtos.UserRegister
	file, err := excelize.OpenReader(buffer)
	if err != nil {
		return usersdot, err
	}

	sheetName := "Sheet1"

	rows, err := file.GetRows(sheetName)

	if err != nil {
		return usersdot, err
	}
	header := rows[0]
	targetCols := []string{"Username", "Name", "Password", "UserLeval"}
	var username []string
	var name []string
	var password []string
	var userlevel []string
	var checked bool
	for _, targetColName := range targetCols {
		checked = false
		for i, colName := range header {
			if colName == targetColName {
				switch targetColName {
				case "Username":
					for _, row := range rows[1:] {
						username = append(username, row[i])
					}
				case "Name":
					for _, row := range rows[1:] {
						name = append(name, row[i])
					}
				case "Password":
					for _, row := range rows[1:] {
						password = append(password, row[i])
					}
				case "UserLeval":
					for _, row := range rows[1:] {
						userlevel = append(userlevel, row[i])
					}
				}
				checked = true
			}
		}
		if !checked {
			if username == nil || name == nil || password == nil || userlevel == nil {
				return usersdot, errors.New("empty data")
			}
		}
	}
	if len(username) == len(name) && len(name) == len(password) && len(password) == len(userlevel) {
		var userdot dtos.UserRegister
		for i := range name {
			userdot.Username = username[i]
			userdot.Name = name[i]
			userdot.Password = password[i]
			userdot.UserLeval = userlevel[i]
			usersdot = append(usersdot, userdot)
		}
	}
	return usersdot, nil
}
