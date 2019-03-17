package pkg

import (
	"encoding/base64"
	"fmt"
	"github.com/mchirico/go_slicestore/yamlpkg"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteFile(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		str := strings.Trim(r.Header.Get("Authorization"), "Basic ")

		data, err := base64.StdEncoding.DecodeString(string(str))
		if err != nil {
			log.Printf("Error in DecodeString and trim: %v", err)
		}

		e := "Mouse:password123"
		if string(data) != e {
			t.Errorf("got header %q, want %q", data, e)
		}

		data, err = ioutil.ReadFile("../fixtures/data")
		if err != nil {

			t.Fatalf("Error ioutil.ReadFile")
		}

		fmt.Fprintln(w, string(data))
	}))
	defer server.Close()

	mockYaml := func() yamlpkg.Config {
		y := yamlpkg.Config{}
		y.Yaml.IP = server.URL
		y.Yaml.Password = "password123"
		y.Yaml.Username = "Mouse"

		return y
	}

	q, err := NewQuery(OptionURL(MockURL), OptionYaml(mockYaml))
	if err != nil {
		t.Fatalf("error")
	}

	dt := DateType{}
	dt.Year = 2018
	dt.Month = 7
	dt.Day = 1
	dt.AddDaysToList(1)
	q.QueryList(dt)

	resultFile, err := ioutil.ReadFile("allData.csv")
	if err != nil {

		t.Fatalf("Error ioutil.ReadFile")
	}

	if len(resultFile) != 1550 {
		t.Fatalf("File not correct")
	}

	expected := "01-01-2019,20922538141144,15682982927025,15164138354,1461714289,107012615940,"
	if strings.Contains(string(resultFile), expected) != true {
		t.Fatalf("Values not in file")
	}

}
