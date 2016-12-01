package main

import (
	"github.com/meapex/meapex/server"
	"github.com/meapex/meapex/server/db"
	"github.com/meapex/meapex/server/models"
	"github.com/pelletier/go-toml"
)

func main() {
	config, _ := toml.LoadFile("./config/config.toml")

	databaseHost := config.Get("database.host").(string)
	databaseDBname := config.Get("database.dbname").(string)
	databaseUsername := config.Get("database.username").(string)
	databasePassword := config.Get("database.password").(string)

	db.InitDB(databaseHost, databaseUsername, databasePassword, databaseDBname)
	db.ORM.AutoMigrate(&models.User{})

	redisHost := config.Get("redis.host").(string)
	redisPort := config.Get("redis.port").(string)
	redisDBname := config.Get("redis.dbname").(int64)
	redisPassword := config.Get("redis.password").(string)

	db.InitPool(redisHost, redisPort, redisPassword, redisDBname)

	server.Init()
}
