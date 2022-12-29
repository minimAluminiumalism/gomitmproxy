package proxyutil

import (
	"io/ioutil"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type Header struct {
	UserAgent string `yaml:"user-agent"`
	Cookie    string `yaml:"cookie"`
}

func ParseConfig(configPath string) (*Header, error) {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	headerConfig := new(Header)
	err = yaml.UnmarshalStrict(yamlFile, headerConfig)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(headerConfig)
	if err != nil {
		return nil, err
	}
	header := new(Header)
	err = json.Unmarshal(data, header)
	if err != nil {
		return nil, err
	}
	return header, nil
}
