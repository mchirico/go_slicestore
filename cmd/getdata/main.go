package main

import (
	"crypto/tls"
	"fmt"
	"github.com/mchirico/go_slicestore/pkg"
	"github.com/mchirico/go_slicestore/yamlpkg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	d := pkg.DS{}
	d.GetStruct(basicAuth())
	fmt.Printf("%v", d.ResponseData.VaultUsageReport.Vaults[0].Name)

}

func basicAuth() []byte {

	y := yamlpkg.Config{}
	file := fmt.Sprintf("%s/sliceStore.yaml", os.Getenv("HOME"))
	y.Read(file)

	url := fmt.Sprintf(
		"https://%s/manager/api/json/1.0/vaultUsageReport.adm?dateRange=true&startDate=03/01/2019&endDate=03/02/2019",
		y.Yaml.IP)

	var username string = y.Yaml.Username
	var passwd string = y.Yaml.Password
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)

	return bodyText
}
