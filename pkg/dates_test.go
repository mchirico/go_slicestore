package pkg

import (
	"log"
	"testing"
)

// We'll need to sort. Ideas: https://play.golang.org/p/KnrxaZqPGwb

func TestTargetDate(t *testing.T) {

	d := DateType{}
	d.Year = 2019
	d.Month = 3
	d.Day = 1

	r := d.TargetDate(6)
	expected := "03-07-2019"

	if r != expected {
		t.Fatalf("Got: %s\nExpected: %s\n", r, expected)
	}

}

func TestStartDate(t *testing.T) {

	d := DateType{}
	d.Year = 2019
	d.Month = 3
	d.Day = 1

	r := d.TargetDate(6)
	expected := "03-07-2019"

	if r != expected {
		t.Fatalf("Got: %s\nExpected: %s\n", r, expected)
	}

}

func Test(t *testing.T) {
	d := DateType{}
	d.Year = 2019
	d.Month = 3
	d.Day = 1


}

func TestDateType_AddDaysToList(t *testing.T) {
	d := DateType{}
	d.Year = 2019
	d.Month = 3
	d.Day = 1
	d.AddDaysToList(3)
	log.Printf("%v\n",d.list.Days)

}


func TestAddDays(t *testing.T) {
	d := DateType{}
	days := Days{}
	day := Day{}

	day.Start = "0"
	day.End = "1"
	days.Days = append(days.Days, day )

	day.Start = "1"
	day.End = "2"
	days.Days = append(days.Days, day )

	d.list = days
	log.Printf("%v\n",d.list.Days)

}



