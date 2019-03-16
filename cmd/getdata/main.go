package main

import (
	"fmt"
	"github.com/mchirico/go_slicestore/pkg"
)

func main() {
	d := pkg.DS{}
	q := pkg.Query{}
	d.GetStruct(q.Query())
	fmt.Printf("%v", d.ResponseData.VaultUsageReport.Vaults[0].Name)


	// Get a list for fixture to test
	pkg.DateFormat = "01/02/2006"
	dt := pkg.DateType{}

	dt.Year = 2019
	dt.Month = 1
	dt.Day = 1
	dt.AddDaysToList(10)
	q.QueryList(dt)

}

