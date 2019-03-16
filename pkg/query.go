package pkg

import (
	"fmt"
	"github.com/mchirico/go_slicestore/yamlpkg"
	"os"
)

type Query struct {}

func (q *Query) Query() [] byte {
	y := yamlpkg.Config{}
	file := fmt.Sprintf("%s/sliceStore.yaml", os.Getenv("HOME"))
	y.Read(file)

	url := fmt.Sprintf(
		"https://%s/manager/api/json/1.0/vaultUsageReport.adm?dateRange=true&startDate=03/01/2019&endDate=03/02/2019",
		y.Yaml.IP)

	return Get(url, y.Yaml.Username, y.Yaml.Password)
}