package excelhelper

import (
	"backend/pkg/variable"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func SetCell(f *excelize.File, sheetName string, colIndex *int, rowIndex int, value interface{}, styleId ...int) error {
	col := variable.IntToAlphabet(*colIndex)
	cell := fmt.Sprintf("%s%d", col, rowIndex)
	if len(styleId) > 0 {
		err := f.SetCellStyle(sheetName, cell, cell, styleId[0])
		if err != nil {
			return err
		}
	}
	defer func() {
		*colIndex++
	}()
	return f.SetCellValue(sheetName, cell, value)
}
