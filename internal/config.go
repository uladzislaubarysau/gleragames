package internal

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Url            string   `json:"url"`
	DataSourceName string   `json:"data_source_name"`
	DBQueryTimeout Duration `json:"db_query_timeout"`
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

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}
