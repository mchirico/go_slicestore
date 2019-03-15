package yamlpkg

// REF: http://sweetohm.net/article/go-yaml-parsers.en.html

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

type y struct {
	IP       string `yaml:"ip"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Config entry point
type Config struct {
	sync.Mutex
	Yaml y
}

// Write yaml file
func (c *Config) Write(file string) error {
	c.Lock()
	defer c.Unlock()

	data, err := yaml.Marshal(c.Yaml)
	if err != nil {
		log.Printf("yaml.Marshal(config): %v", err)
		return err
	}

	err = ioutil.WriteFile(file, data, 0600)
	if err != nil {
		log.Printf("error in yaml write: %v", err)
	}
	return err
}

// Read yaml file
func (c *Config) Read(file string) error {
	c.Lock()
	defer c.Unlock()

	source, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Error ioutil.ReadFile")
		return err
	}
	err = yaml.Unmarshal(source, &c.Yaml)
	if err != nil {
		log.Printf("Error Unmarshal")
		return err
	}

	return err
}

// SetDefault simple config settings
func (c *Config) SetDefault() {
	c.Lock()
	defer c.Unlock()
	c.Yaml.IP = "0.0.0.0"
	c.Yaml.Username = "Spock"
	c.Yaml.Password = "Password223"

}
