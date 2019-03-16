package pkg

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var p pdata

func TestMain(m *testing.M) {
	SetupFunction()
	retCode := m.Run()
	TeardownFunction()
	os.Exit(retCode)
}

func SetupFunction() {

	p = Pdata()
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

func TestPdata_FormatNames(t *testing.T) {

	p.FormatNames()
	count := 0
	list0 := []string{}
	list1 := []string{}

	for _, v := range p.list {
		list0 = append(list0, p.formatname[v])
		list1 = append(list1, v)
		count++
	}

	if list0[62] != "stage_ipvod17" {
		t.Fatalf("Expected: %s\nGot:%s\n",
			"stage_ipvod17", list0[62])

	}
	if list1[62] != "stage_ipvod000017" {
		t.Fatalf("Expected: %s\nGot:%s\n",
			"stage_ipvod000017", list1[62])

	}

}

func TestPdata_Write(t *testing.T) {
	p.FormatNames()
	err := p.Write("junk.csv")
	if err != nil {
		t.Fatalf("error in write: %v", err)
	}
}
