package pkg

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
)




func Get(url string, username string, password string) []byte {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error in Get: %v\n", err)
	}

	return bodyText

}
