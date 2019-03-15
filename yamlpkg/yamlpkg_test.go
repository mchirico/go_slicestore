package yamlpkg

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestReadError_File_Problem(t *testing.T) {

	var str bytes.Buffer
	log.SetOutput(&str)

	c := Config{}
	err := c.Read("///")
	if err == nil {
		t.FailNow()
	}
	if strings.Contains(str.String(),
		"Error ioutil.ReadFile") != true {
		t.Fatalf("Got:\n%s\n\nExpected:\n%s\n", str.String(),
			"Error ioutil.ReadFile")
	}

}

func TestReadError_Content_Problem(t *testing.T) {

	file := "_junk.yaml"

	d1 := []byte("a\n\n\b\n")
	err := ioutil.WriteFile(file, d1, 0644)

	var str bytes.Buffer
	log.SetOutput(&str)

	c := Config{}
	err = c.Read(file)
	if err == nil {
		t.FailNow()
	}
	if strings.Contains(str.String(),
		"Error Unmarshal") != true {
		t.Fatalf("Got:\n%s\n\nExpected:\n%s\n", str.String(),
			"Error Unmarshal")
	}

}

func TestWriteError_File_Problem(t *testing.T) {

	file := "///"
	var str bytes.Buffer
	log.SetOutput(&str)

	c := Config{}
	err := c.Write(file)
	if err == nil {
		t.FailNow()
	}

	expected := `error in yaml write: open ///: is a directory`

	if strings.Contains(str.String(),
		expected) != true {
		t.Fatalf("Got:\n%s\n\nExpected:\n%s\n", str.String(),
			expected)
	}

}

func TestReadWrite(t *testing.T) {
	file := "sliceStore.yaml"

	c := Config{}

	c.Yaml.IP = "0.0.0.0"
	c.Yaml.Username = "Tom"
	c.Yaml.Password = "pass2lks"

	c.Write(file)

	c2 := Config{}
	c2.Read(file)

	if c2.Yaml.Username != c.Yaml.Username {
		t.Fatalf("Got:\n%s\n\nExpected:\n%s\n", c2.Yaml.Username, c.Yaml.Username)
	}

}

func TestConfig_SetDefault(t *testing.T) {
	c := Config{}
	c.SetDefault()
	if c.Yaml.Username != "Spock" {
		t.FailNow()
	}

}
