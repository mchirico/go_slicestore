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

func mockURL(ip string, start string, end string) string {
	log.Printf("https://%s/manager/api/json/1.0/"+
		"vaultUsageReport.adm?"+
		"dateRange=true&startDate=%s&endDate=%s", ip, start, end)

	return ip
}

func optionURL(t URLtype) func(f *Query) error {
	return func(f *Query) error {
		f.urlFunction = t
		return nil
	}
}

func optionYaml(t YamlFunction) func(f *Query) error {
	return func(f *Query) error {
		f.yamlconfig = t
		return nil
	}
}

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

		log.Printf("----------------\n\n*************\n")

		fmt.Fprintln(w, string(data))
	}))
	defer server.Close()

	//result := Get(server.URL, "Mouse", "password123")
	//log.Printf("result: %v\n", string(result))

	mockYaml := func() yamlpkg.Config {
		y := yamlpkg.Config{}
		y.Yaml.IP = server.URL
		y.Yaml.Password = "password123"
		y.Yaml.Username = "Mouse"

		return y
	}

	q, err := NewQuery(optionURL(mockURL), optionYaml(mockYaml))
	if err != nil {
		t.Fatalf("error")
	}

	dt := DateType{}
	dt.Year = 2018
	dt.Month = 7
	dt.Day = 1
	dt.AddDaysToList(1)
	q.QueryList(dt)

}
