package pkg

import (
	"io/ioutil"
	"testing"
)

func TestWriteFile(t *testing.T) {

	data := []byte("this is a test")
	ioutil.WriteFile("/tmp/dat1", data, 0644)

}
