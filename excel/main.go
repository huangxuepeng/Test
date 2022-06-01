package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("cc")
	if err != nil {
		panic(err.Error())
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "价格"
	for i := len(arr) - 1; i >= 0; i-- {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = fmt.Sprintln(arr[i])
	}

	err = file.Save("cc.xlsx")
	if err != nil {
		panic(err.Error())
	}
}
