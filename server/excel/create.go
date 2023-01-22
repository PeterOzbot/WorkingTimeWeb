package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func Create(requests []*Request) *excelize.File {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			//fmt.Println(err)
		}
	}()

	// // Create a new sheet.
	// _, err := f.NewSheet("Sheet1")
	// if err != nil {
	// 	//fmt.Println(err)
	// }

	// style, _ := f.NewStyle(&excelize.Style{
	// 	for
	// })
	//f.SetColStyle("Sheet1","A",style)

	f.SetCellValue("Sheet1", "A1", "Datum")
	f.SetColWidth("Sheet1", "A", "B", 15)

	f.SetCellValue("Sheet1", "B1", "Ure")
	f.SetColWidth("Sheet1", "B", "C", 10)

	f.SetCellValue("Sheet1", "C1", "Opis")
	f.SetColWidth("Sheet1", "C", "D", 50)

	finalIndex := 0
	for index, request := range requests {

		dateCell := fmt.Sprintf("A%d", index+2)
		hoursCell := fmt.Sprintf("B%d", index+2)

		f.SetCellValue("Sheet1", dateCell, request.Date.Format("02.01.2006"))
		f.SetCellValue("Sheet1", hoursCell, request.Hours)

		finalIndex = index + 2
	}

	hoursSumTextCell := fmt.Sprintf("A%d", finalIndex+2)
	f.SetCellValue("Sheet1", hoursSumTextCell, "Skupaj")

	hoursSumCell := fmt.Sprintf("B%d", finalIndex+2)
	formula := fmt.Sprintf("SUM(B2:B%d)", finalIndex+1)
	f.SetCellFormula("Sheet1", hoursSumCell, formula)

	return f

	//return strconv.Itoa(len(request))
}
