package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ORM *gorm.DB
)

func InitDB(host string, username string, password string, dbname string) *gorm.DB {
	var err error
	ORM, err = gorm.Open("postgres", "host="+host+" user="+username+" password="+password+" dbname="+dbname+" sslmode=disable")
	if err != nil {
		panic(err)
		return nil
	}

	return ORM
}
