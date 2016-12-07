package server

import (
	"github.com/gin-gonic/gin"
	"github.com/meapex/meapex/server/controllers"
	"github.com/meapex/meapex/server/controllers/api"
	"github.com/meapex/meapex/server/db"
	"github.com/meapex/meapex/server/models"
	"github.com/pelletier/go-toml"
)

// Init ...
func Init() {
	r := gin.Default()

	initDB()

	account := r.Group("/account")
	{
		account.POST("/signup", controllers.Signup)
		account.POST("/signin", controllers.Signin)
		account.POST("/signout", controllers.Signout)
		account.POST("/find-password", controllers.FindPassword)
	}

	apiv1 := r.Group("/api")
	{
		apiv1.GET("/me", api.Me)
		apiv1.GET("/resources", api.GetAllResource)
		apiv1.GET("/resources/:id", api.GetResourceById)
		apiv1.POST("/resources", api.CreateResource)
	}

	r.Run(":8080")
}

// initDB ...
func initDB() {
	config, _ := toml.LoadFile("./config/config.toml")

	databaseHost := config.Get("database.host").(string)
	databasePort := config.Get("database.port").(string)
	databaseDBname := config.Get("database.dbname").(string)
	databaseUsername := config.Get("database.username").(string)
	databasePassword := config.Get("database.password").(string)

	db.InitDB(databaseHost, databasePort, databaseUsername, databasePassword, databaseDBname)
	db.ORM.AutoMigrate(&models.User{}, &models.Resource{})

	redisHost := config.Get("redis.host").(string)
	redisPort := config.Get("redis.port").(string)
	redisDBname := config.Get("redis.dbname").(int64)
	redisPassword := config.Get("redis.password").(string)

	db.InitRedisPool(redisHost, redisPort, redisPassword, redisDBname)
}
