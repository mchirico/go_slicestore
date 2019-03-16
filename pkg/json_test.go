package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

// We'll need to sort. Ideas: https://play.golang.org/p/KnrxaZqPGwb

func TestRead(t *testing.T) {
	source, err := ioutil.ReadFile("../fixtures/data")
	if err != nil {

		t.Fatalf("Error ioutil.ReadFile")
	}
	//log.Printf("%s\n", source)

	s := DS{}
	err = json.Unmarshal(source, &s)
	if err != nil {
		t.Fatalf("Can't Unmarshal")
	}

	fmt.Printf("%v", s.ResponseData.VaultUsageReport.Vaults[0].Name)

}

func TestGetStruct(t *testing.T) {
	source, err := ioutil.ReadFile("../fixtures/data")
	if err != nil {

		t.Fatalf("Error ioutil.ReadFile")
	}
	//log.Printf("%s\n", source)

	d := DS{}
	d.GetStruct(source)

	fmt.Printf("%v", d.ResponseData.VaultUsageReport.Vaults[0].Name)

}

func TestWorkWithDecoded(t *testing.T) {
	source, err := ioutil.ReadFile("../fixtures/data")
	if err != nil {

		t.Fatalf("Error ioutil.ReadFile")
	}
	//log.Printf("%s\n", source)

	d := DS{}
	d.GetStruct(source)

	start := d.ResponseData.VaultUsageReport.DateRange.StartDate
	if start != "2019-01-01" {
		t.Fatalf("Can not read start:\n%s\n", start)
	}

	length := len(d.ResponseData.VaultUsageReport.Vaults)
	if length != 90 {
		t.Fatalf("Length not correct:\n%d\n", length)
	}

}
