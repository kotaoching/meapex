package main

import (
	"github.com/meapex/meapex/server"
	"github.com/meapex/meapex/server/models"
	"github.com/pelletier/go-toml"
)

func main() {
	config, _ := toml.LoadFile("./config/config.toml")
	host := config.Get("database.host").(string)
	database := config.Get("database.database").(string)
	username := config.Get("database.username").(string)
	password := config.Get("database.password").(string)

	db := models.InitDB("host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable")
	db.AutoMigrate(&models.User{})

	server.Init()
}
