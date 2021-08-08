package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type yamlConfig struct {
	Servers []struct {
		Server    string `yaml:"server"`
		Name      string `yaml:"name"`
		Databases []struct {
			Database string `yaml:"database"`
			Name     string `yaml:"name"`
			Tables   []struct {
				Table string `yaml:"table"`
				Name  string `yaml:"name"`
			} `yaml:"tables"`
		} `yaml:"databases"`
	} `yaml:"servers"`
}

var cfg yamlConfig

func init() {

	yamlConfig, err := ioutil.ReadFile(".config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(yamlConfig, &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	//	fmt.Println(cfg)
}

func main() {
	//	fmt.Println(cfg.Servers[0])
	for _, v := range cfg.Servers {
		fmt.Println(v)
	}
}
