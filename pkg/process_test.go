package pkg

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var p  Pdata


func TestMain(m *testing.M) {
	SetupFunction()
	retCode := m.Run()
	TeardownFunction()
	os.Exit(retCode)
}

func SetupFunction() {

	p = Pdata{}
	data, err := ioutil.ReadFile("../fixtures/" +
		"fixtures_01_01_2019_01_02_2019")

	data2, err := ioutil.ReadFile("../fixtures/" +
		"fixtures_01_02_2019_01_03_2019")

	if err != nil {

		log.Fatalf("Error ioutil.ReadFile")
	}
	p.Add(data)
	p.Add(data2)

}

func TeardownFunction() {

}


func TestProcess(t *testing.T) {

	if len(p.raw[0].ResponseData.VaultUsageReport.Vaults) != 90 {
		t.Fatalf("Problem with read. Expected 90")
	}

	if len(p.raw) != 2 {
		t.Fatalf("Problem with p.raw  Expected 2")
	}

}

func TestProcess2(t *testing.T) {

	if len(p.raw[0].ResponseData.VaultUsageReport.Vaults) != 90 {
		t.Fatalf("Problem with read. Expected 90")
	}

	if len(p.raw) != 2 {
		t.Fatalf("Problem with p.raw  Expected 2")
	}

}








