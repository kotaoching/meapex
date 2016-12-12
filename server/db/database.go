package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmoiron/sqlx"
)

var (
	ORM  *gorm.DB
	Sqlx *sqlx.DB
)

func InitORM(host string, port string, username string, password string, dbname string) *gorm.DB {
	var err error
	ORM, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" password="+password+" dbname="+dbname+" sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return ORM
}

func InitSqlx(host string, port string, username string, password string, dbname string) *sqlx.DB {
	var err error

	Sqlx, err = sqlx.Connect("postgres", "host="+host+" port="+port+" user="+username+" password="+password+" dbname="+dbname+" sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return Sqlx
}
