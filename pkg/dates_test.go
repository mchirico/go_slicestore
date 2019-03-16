package pkg

import (
	"log"
	"reflect"
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

	DateFormat = "01-02-2006"

	d := DateType{}
	d.Year = 2019
	d.Month = 3
	d.Day = 1
	d.AddDaysToList(3)
	expected := []Day{{"03-01-2019", "03-02-2019"},
		{"03-02-2019", "03-03-2019"},
		{"03-03-2019", "03-04-2019"},
	}

	diff := reflect.DeepEqual(d.list.Days, expected)

	if diff != true {
		t.Fatalf("Expected: %v\nGot: %v\n", expected, d.list.Days)
	}
	log.Printf("%v\n", d.list.Days)

}

func TestDateType_SubtractDaysToList(t *testing.T) {

	DateFormat = "01-02-2006"

	d := DateType{}
	d.Year = 2019
	d.Month = 3
	d.Day = 1
	d.AddDaysToList(-3)
	expected := []Day{{"03-01-2019", "02-28-2019"},
		{"02-28-2019", "02-27-2019"},
		{"02-27-2019", "02-26-2019"},
	}

	diff := reflect.DeepEqual(d.list.Days, expected)

	if diff != true {
		t.Fatalf("Expected: %v\nGot: %v\n", expected, d.list.Days)
	}
	log.Printf("%v\n", d.list.Days)

}

func TestAddDays(t *testing.T) {
	d := DateType{}
	days := Days{}
	day := Day{}

	day.Start = "0"
	day.End = "1"
	days.Days = append(days.Days, day)

	day.Start = "1"
	day.End = "2"
	days.Days = append(days.Days, day)

	d.list = days

	expected := []Day{{"0", "1"}, {"1", "2"}}

	diff := reflect.DeepEqual(d.list.Days, expected)

	if diff != true {
		t.Fatalf("Expected: %v\nGot: %v\n", expected, d.list.Days)
	}

}

func TestDateType_Format(t *testing.T) {

	DateFormat = "01/02/2006"

	d := DateType{}
	d.Year = 2019
	d.Month = 3
	d.Day = 1
	d.AddDaysToList(3)
	expected := []Day{{"03/01/2019", "03/02/2019"},
		{"03/02/2019", "03/03/2019"},
		{"03/03/2019", "03/04/2019"},
	}

	diff := reflect.DeepEqual(d.list.Days, expected)

	if diff != true {
		t.Fatalf("Expected: %v\nGot: %v\n", expected, d.list.Days)
	}

	log.Printf("%v\n", d.list.Days)

}
