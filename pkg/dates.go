package pkg

import (
	"sync"
	"time"
)

var DateFormat = "01-02-2006"

type Day struct {
	Start string
	End   string
}

type Days struct {
	Days []Day
}

type DateType struct {
	Year  int
	Month int
	Day   int
	list  Days
	sync.Mutex
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (d *DateType) TargetDate(days int) string {
	return date(d.Year, d.Month, d.Day).AddDate(0, 0, days).Format(DateFormat)
}

func (d *DateType) AddDaysToList(num int) {
	d.Lock()
	defer d.Unlock()

	if num < 0 {
		d.subDaysToList(num)
		return
	}

	days := Days{}
	day := Day{}
	for i := 1; i <= num; i++ {
		day.Start = d.TargetDate(i - 1)
		day.End = d.TargetDate(i)
		days.Days = append(days.Days, day)

	}

	d.list = days

}

func (d *DateType) subDaysToList(num int) {

	days := Days{}
	day := Day{}
	for i := -1; i >= num; i-- {
		day.Start = d.TargetDate(i + 1)
		day.End = d.TargetDate(i)
		days.Days = append(days.Days, day)

	}

	d.list = days

}

func (d *DateType) Start() string {
	return date(d.Year, d.Month, d.Day).Format(DateFormat)
}
