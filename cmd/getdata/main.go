package main

import (
	"github.com/mchirico/go_slicestore/pkg"
)

func main() {

	q := pkg.Query{}
	pkg.DateFormat = "01/02/2006"
	dt := pkg.DateType{}

	dt.Year = 2018
	dt.Month = 7
	dt.Day = 1
	dt.AddDaysToList(258)
	q.QueryList(dt)

}
