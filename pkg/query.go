package pkg

import (
	"fmt"
	"github.com/mchirico/go_slicestore/yamlpkg"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type URLtype func(string, string, string) string
type YamlFunction func() yamlpkg.Config
type YamlValue yamlpkg.Config

type Query struct {
	yamlconfig  YamlFunction
	urlFunction URLtype
}

func NewQuery(options ...func(*Query) error) (*Query, error) {
	f := &Query{}

	f.yamlconfig = defaultYaml
	f.urlFunction = defaultURL

	for _, op := range options {
		err := op(f)
		if err != nil {
			return nil, err
		}
	}

	return f, nil
}

func defaultYaml() yamlpkg.Config {
	y := yamlpkg.Config{}
	file := fmt.Sprintf("%s/sliceStore.yaml", os.Getenv("HOME"))
	y.Read(file)
	return y
}

func defaultURL(ip string, start string, end string) string {
	url := fmt.Sprintf("https://%s/manager/api/json/1.0/"+
		"vaultUsageReport.adm?"+
		"dateRange=true&startDate=%s&endDate=%s", ip, start, end)
	return url
}

func (q *Query) QueryList(d DateType) {

	y := q.yamlconfig()
	p := Pdata()

	for _, v := range d.list.Days {

		url := q.urlFunction(y.Yaml.IP, v.Start, v.End)
		data := Get(url, y.Yaml.Username, y.Yaml.Password)

		vstart := strings.Replace(v.Start, "/", "_", -1)
		vend := strings.Replace(v.End, "/", "_", -1)

		file := fmt.Sprintf("fixtures_%s_%s", vstart, vend)
		log.Printf("write: %s\n", file)

		p.Add(data)

	}
	p.FormatNames()
	p.Write("allData.csv")
}

func Write(file string, data []byte) error {
	return ioutil.WriteFile(file, data, 0644)
}

/* +++++++++++++++++++++++++++++++++++++++++++++++++++++
                      For testing
   +++++++++++++++++++++++++++++++++++++++++++++++++++++

Reference in query_test.go the following:

mockYaml := func() yamlpkg.Config {
		y := yamlpkg.Config{}
		y.Yaml.IP = server.URL
		y.Yaml.Password = "password123"
		y.Yaml.Username = "Mouse"

		return y
	}



++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */

func MockURL(ip string, start string, end string) string {

	debug := false

	if debug {
		log.Printf("Example call:\n%s/manager/api/json/1.0/"+
			"vaultUsageReport.adm?"+
			"dateRange=true&startDate=%s&endDate=%s\n", ip, start, end)
	}
	
	return ip
}

func OptionURL(t URLtype) func(f *Query) error {
	return func(f *Query) error {
		f.urlFunction = t
		return nil
	}
}

func OptionYaml(t YamlFunction) func(f *Query) error {
	return func(f *Query) error {
		f.yamlconfig = t
		return nil
	}
}
