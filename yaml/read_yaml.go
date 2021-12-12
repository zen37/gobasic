package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
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

	Upper1 struct {
		Name string `yaml:"name"`
	}

	Upper2 struct {
		Name string `yaml:"name"`
	} `yaml:"UPPER"`

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

type ignoreRequest struct {
	Prefix    string `yaml:"prefix"`
	OtherName string `yaml:"suffix"`
	Regex     string
}

var (
	cfg    yamlConfig
	cfg2   yamlConfig2
	cfgReq ignoreRequest
)

func init() {

	yamlConfig, err := ioutil.ReadFile(".config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(yamlConfig, &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("----- nested config -----")
	fmt.Println(cfg)

	yamlConfig2, err := ioutil.ReadFile(".config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(yamlConfig2, &cfg2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("----- struct naming for yaml -----")
	fmt.Println(cfg2)

	yamlConfigReq, err := ioutil.ReadFile(".config2.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	//data := make(map[string]interface{})
	//data := make([]string, 10)
	err = yaml.Unmarshal(yamlConfigReq, &cfgReq)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("----- ignore requests  -----")
	fmt.Println("prefix:", cfgReq.Prefix)
	fmt.Println("other name:", cfgReq.OtherName)
	fmt.Println("regex:", cfgReq.Regex)
	fmt.Println("regex 1st:", strings.Fields(cfgReq.Regex)[0])

}

func main() {
	//	fmt.Println(cfg.Servers[0])
	//for _, v := range cfg.Servers {
	//fmt.Println(v)
	//}
}
