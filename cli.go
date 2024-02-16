package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/chadmiller/energyratecalculation/tariffs"

	"github.com/xuri/excelize/v2"
)

func interpretFile(filename string) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		fmt.Println(filename)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
			fmt.Println(filename)
		}
	}()

	calculators := tariffs.GetCalculators("usa/sc/ga")

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		fmt.Println(filename)
		return
	}
	for i, row := range rows {
		if len(row) != 3 {
			fmt.Printf("error parsing file at line %d: %v\n", i, row)
			continue
		}
		if i == 0 && row[0] == "Hour" {
			continue
		}
		datetime, err := time.Parse("2006-01-02 15:04", row[0])
		if err != nil {
			fmt.Printf("error parsing file at column 1, line %d: %v\n", i+1, row[0])
			continue
		}
		kWh, err := strconv.ParseFloat(row[1], 32)
		if err != nil {
			fmt.Printf("error parsing file at column 2, line %d: %v\n", i+1, row[1])
			continue
		}
		tempf, err := strconv.ParseFloat(row[2], 32)
		if err != nil {
			fmt.Printf("error parsing file at column 3, line %d: %v\n", i+1, row[2])
			continue
		}

		tariffs.AddRowToAll(calculators, datetime, tariffs.UsagekWh(kWh), tariffs.TemperatureFahrenheit(tempf))
	}

	fmt.Printf("from file %v\n", filename)
	for _, calculator := range(calculators) {
		fmt.Printf("%20s: $%0.2f  %s\n", calculator, float32(calculator.Compute() / 100.0), calculator.Describe())
	}
}

func main() {
	flag.Parse()
	for i := 1; i < len(os.Args); i++ {
		interpretFile(os.Args[i])
	}
}

// vi: set shiftwidth=4 tabstop=4 noexpandtab nolist :
