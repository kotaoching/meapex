package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ORM *gorm.DB
)

func InitDB(dataSourceName string) *gorm.DB {
	var err error
	ORM, err = gorm.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
		return nil
	}

	return ORM
}
