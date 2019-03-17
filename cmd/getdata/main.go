package main

import (
	"github.com/mchirico/go_slicestore/pkg"
	"log"
)

func main() {

	q, err := pkg.NewQuery()
	if err != nil {
		log.Fatalf("Error reading pkg.NewQuery()")
	}
	pkg.DateFormat = "01/02/2006"
	dt := pkg.DateType{}

	dt.Year = 2018
	dt.Month = 7
	dt.Day = 1
	dt.AddDaysToList(25)
	q.QueryList(dt)

}
