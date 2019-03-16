package pkg

import (
	"io/ioutil"
	"testing"
)

func TestProcess(t *testing.T) {
	data, err := ioutil.ReadFile("../fixtures/" +
		"fixtures_01_01_2019_01_02_2019")

	if err != nil {

		t.Fatalf("Error ioutil.ReadFile")
	}

	s, err := Process(data)
	if err != nil {
		t.FailNow()
	}

	if len(s.ResponseData.VaultUsageReport.Vaults) != 90 {
		t.Fatalf("Problem with read. Expected 90")
	}


}









