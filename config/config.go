package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Secrets struct {
	CoinMarketCapApiKey string `yaml:"coinmarketcap-api-key"`
}

func LoadSecrets(secrestsFile string) (*Secrets, error) {
	yamlFile, err := ioutil.ReadFile(secrestsFile)
	if err != nil {
		return nil, err
	}
	secrets := &Secrets{}
	err = yaml.Unmarshal(yamlFile, secrets)
	return secrets, err
}
