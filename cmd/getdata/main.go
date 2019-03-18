/*
MIT License

Copyright (c) 2019 Mike Chirico

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package main

import (
	"flag"
	"fmt"
	"github.com/mchirico/date/parse"
	"github.com/mchirico/go_slicestore/pkg"
	"log"
	"time"
)

var startDateFlag = flag.String("start", "6/1/2018", "Start Date")
var endDateFlag = flag.String("end", "now", "End Date")

func init() {
	flag.Usage = func() {
		fmt.Printf(`


Examples:

    	  ./getdata -start "1/1/2019"




Note:

  You must have a ~/sliceStore.yaml setup with your information.


cat > ~/sliceStore.yaml <<EOF
ip: 10.10.10.10
username: Spock
password: Password123
EOF



`)

		//flag.PrintDefaults()
	}
	flag.Parse()

}

func process(q *pkg.Query) {

	pkg.DateFormat = "01/02/2006"
	dt := pkg.DateType{}

	tt, err := parse.DateTimeParse(*startDateFlag).GetTimeLoc()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	days := int(time.Since(tt).Hours() / 24)

	dt.Year = int(tt.Year())
	dt.Month = int(tt.Month())
	dt.Day = int(tt.Day())
	dt.AddDaysToList(days + 1)

	fmt.Printf("Year: %d  Month: %d  Day: %d  AddDays: %d\n",
		dt.Year, dt.Month, dt.Day, days+1)

	q.QueryList(dt)

}

func main() {

	q, err := pkg.NewQuery()
	if err != nil {
		log.Fatalf("Error reading pkg.NewQuery()")
	}
	process(q)

}
