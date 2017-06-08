package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config structure representing the conf used by the application
type Config struct {
	DBName     string
	DBUrl      string
	DBUser     string
	DBPassword string
	Port       int
}

// LoadConfig read from file and unmarshall to Config structure.
func LoadConfig(configFile string) (*Config, error) {
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	if err = json.Unmarshal(conf, c); err != nil {
		return nil, err
	}
	return c, nil
}
