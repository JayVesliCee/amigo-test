/*
	A config file is often use through a whole architecture.
	Thus it could be exported as an other package.
*/

package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config structure representing the conf used by the application
type Config struct {
	DBName     string
	DBURL      string
	DBUser     string
	DBPassword string
	Port       int
}

// LoadConfig read from file and unmarshall to Config structure.
func loadConfig(configFile string) (*Config, error) {
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
