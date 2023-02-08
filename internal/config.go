package internal

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Config used to initialize service
type Config struct {
}

func NewConfig() (*Config, error) {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	configFile := fs.String("config", "./config/config-sample.json", "config file path")
	err := fs.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}
	c := Config{}
	LoadConfigFromFile(*configFile, &c)

	return &c, nil
}

// LoadConfigFromFile parsing json file to config struct
func LoadConfigFromFile(configPath string, c *Config) {
	afp, err := filepath.Abs(configPath)
	if err != nil {
		panic(err)
	}
	jsonConfig, err := os.Open(afp)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := jsonConfig.Close(); err != nil {
			panic(err)
		}
	}()

	data, err := ioutil.ReadAll(jsonConfig)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
}
