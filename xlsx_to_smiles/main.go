package main

import (
	"bufio"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"os"
)

func main() {
	showHelpMsg()

	f, err := excelize.OpenFile("smiles.xlsx")
	check(err)

	err = os.MkdirAll("./smiles", os.ModePerm)
	check(err)

	for i := 1; ; i++ {
		name := f.GetCellValue("Sheet1", fmt.Sprintf("A%d", i))
		smi := f.GetCellValue("Sheet1", fmt.Sprintf("B%d", i))
		if smi == "" || name == "" {
			break
		}

		err := ioutil.WriteFile(fmt.Sprintf("./smiles/%s.smi", name), []byte(smi), 0644)
		check(err)

		fmt.Println(i, "\t",  name, "\t", smi)
	}

	fmt.Print("DONE!:", "Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}


func check(e error) {
	if e != nil {
		panic(e)
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}

func showHelpMsg() {
	fmt.Println("Make sure the following rules are met:")
	fmt.Printf("\t1, Your excel file is named 'smiles'\n")
	fmt.Printf("\t2, The data you want to extract is in a sheet called 'Sheet1'\n")
	fmt.Printf("\t3, The first and second column of your sheet are the names and the smi strings respectively'\n")
	fmt.Printf("\t4, There are no empty rows between data rows. The program assumes an empty row is the end of data\n")
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	fmt.Println("The smi files will be created in a directory named 'smiles'")
	fmt.Println("Any conflicting names will be overwritten")
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
