package helper

import (
	"backend/pkg/https"
	"backend/pkg/logger"
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
)

func UnexpectedError(w http.ResponseWriter, r *http.Request, err error) {
	logger.Trace(err)
	https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
}

func CompareColLargestWidth(columnLargestWidth *[]int, rowDatas []interface{}) {
	for i, d := range rowDatas {
		cellWidth := utf8.RuneCountInString(fmt.Sprintf("%v", d)) + 8 // + 8 for margin
		if len(*columnLargestWidth) < i+1 {
			*columnLargestWidth = append(*columnLargestWidth, cellWidth)
		} else if (*columnLargestWidth)[i] < cellWidth {
			(*columnLargestWidth)[i] = cellWidth
		}
	}
}

func SetColumnWidth(columnWidths *[]int, sheeName string, f *excelize.File) error {
	for i, colWidth := range *columnWidths {
		name, err := excelize.ColumnNumberToName(i + 1)
        if err != nil {
                return err
        }
        f.SetColWidth(sheeName, name, name, float64(colWidth))
	}
	return nil
}