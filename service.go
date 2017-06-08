package main

import (
	"github.com/amigo-test/config"
	"github.com/jinzhu/gorm"
)

type service struct {
	Config *config.Config
	DB     *gorm.DB
}

func newService(configFile string) (*service, error) {
	c, err := config.LoadConfig(configFile)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(
		"postgres",
		"host="+c.DBURL+" user="+c.DBUser+" dbname="+c.DBName+" sslmode=disable password="+c.DBPassword)
	if err != nil {
		return nil, err
	}

	s := &service{
		Config: c,
		DB:     db,
	}

	return s, nil
}
