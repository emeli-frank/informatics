package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"os"
)

func main() {
	f, err := excelize.OpenFile("smiles.xlsx")
	check(err)

	os.MkdirAll("./smiles", os.ModePerm)

	// Get value from cell by given worksheet name and axis.
	for i := 1; ; i++ {
		name := f.GetCellValue("Sheet1", fmt.Sprintf("A%d", i))
		smi := f.GetCellValue("Sheet1", fmt.Sprintf("B%d", i))
		if smi == "" {
			break
		}

		err := ioutil.WriteFile(fmt.Sprintf("./smiles/%s.smi", name), []byte(smi), 0644)
		check(err)

		//fmt.Println(i, "\t",  name, "\t", smi)
	}
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}
