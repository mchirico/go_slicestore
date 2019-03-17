package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/mchirico/go_slicestore/pkg"
	"github.com/mchirico/go_slicestore/yamlpkg"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	SetupFunction()
	retCode := m.Run()
	TeardownFunction()
	os.Exit(retCode)
}

func SetupFunction() {
	flag.Parse()
}

func TeardownFunction() {

}

func Test_process(t *testing.T) {

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

		data, err = ioutil.ReadFile("../../fixtures/mainData")
		if err != nil {

			t.Fatalf("Error ioutil.ReadFile")
		}

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

	q, err := pkg.NewQuery(pkg.OptionURL(pkg.MockURL), pkg.OptionYaml(mockYaml))
	if err != nil {
		t.Fatalf("error")
	}

	process(q)

	resultFile, err := ioutil.ReadFile("allData.csv")
	if err != nil {

		t.Fatalf("Error ioutil.ReadFile")
	}

	if len(resultFile) != 208651 {
		t.Fatalf("File not correct: %v\n", len(resultFile))
	}

	expected := "01-01-2019,2"
	if strings.Contains(string(resultFile), expected) != true {
		t.Fatalf("Values not in file")
	}

}
