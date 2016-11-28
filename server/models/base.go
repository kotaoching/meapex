package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func InitDB(dataSourceName string) *gorm.DB {
	var err error
	db, err = gorm.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
		return nil
	}

	return db
}
