package main

import "github.com/amigo-test/config"

type service struct {
	Config *config.Config
}

func newService(configFile string) (*service, error) {
	c, err := config.LoadConfig(configFile)
	if err != nil {
		return nil, err
	}

	s := &service{
		Config: c,
	}
	return s, nil
}
