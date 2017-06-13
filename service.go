/*
	Service structure is the main manager for our configuration flow.
	We manage the main conf (url, keys, pwd) and the DB session
*/

package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type service struct {
	Config *Config
	DB     *gorm.DB
}

func newService(configFile string) (*service, error) {
	c, err := loadConfig(configFile)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(
		"postgres",
		"host="+c.DBURL+" user="+c.DBUser+" dbname="+c.DBName+" sslmode=disable password="+c.DBPassword)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&message{})

	s := &service{
		Config: c,
		DB:     db,
	}

	return s, nil
}

func (s *service) Close() {
	s.DB.Close()
}
