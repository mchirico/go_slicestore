package pkg

import (
	"fmt"
	"github.com/mchirico/go_slicestore/yamlpkg"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Query struct{}

func (q *Query) QueryList(d DateType) {
	y := yamlpkg.Config{}
	file := fmt.Sprintf("%s/sliceStore.yaml", os.Getenv("HOME"))
	y.Read(file)

	for _, v := range d.list.Days {

		url := fmt.Sprintf("https://%s/manager/api/json/1.0/"+
			"vaultUsageReport.adm?"+
			"dateRange=true&startDate=%s&endDate=%s", y.Yaml.IP, v.Start, v.End)

		data := Get(url, y.Yaml.Username, y.Yaml.Password)

		vstart := strings.Replace(v.Start, "/", "_", -1)
		vend := strings.Replace(v.End, "/", "_", -1)

		file := fmt.Sprintf("fixtures_%s_%s", vstart, vend)
		log.Printf("write: %s\n", file)
		Write(file, data)

	}
}

func (q *Query) Query() []byte {
	y := yamlpkg.Config{}
	file := fmt.Sprintf("%s/sliceStore.yaml", os.Getenv("HOME"))
	y.Read(file)

	url := fmt.Sprintf(
		"https://%s/manager/api/json/1.0/vaultUsageReport.adm?dateRange=true&startDate=03/01/2019&endDate=03/02/2019",
		y.Yaml.IP)

	return Get(url, y.Yaml.Username, y.Yaml.Password)
}

func Write(file string, data []byte) error {
	return ioutil.WriteFile(file, data, 0644)
}
