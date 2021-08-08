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

type yamlConfig2 struct {
	Lower struct {
		Name string `yaml:"name"`
	}

	Upper struct {
		Name string `yaml:"name"`
	}

	camelCase struct {
		Name string `yaml:"name"`
	}

	PascalCase struct {
		Name string `yaml:"name"`
	}

	AnyName struct {
		Name string `yaml:"name"`
	} `yaml:"PascalCase"`

	AnyName2 struct {
		Name string `yaml:"name"`
	} `yaml:"camelCase"`
}

var cfg yamlConfig

var cfg2 yamlConfig2

func init() {

	yamlConfig, err := ioutil.ReadFile(".config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(yamlConfig, &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	yamlConfig2, err := ioutil.ReadFile(".config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(yamlConfig2, &cfg2)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(cfg2)
}

func main() {
	//	fmt.Println(cfg.Servers[0])
	for _, v := range cfg.Servers {
		fmt.Println(v)
	}
}
