package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meapex/meapex/server/controllers"
	"github.com/meapex/meapex/server/controllers/api"
)

func Init() {
	r := gin.Default()

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
		apiv1.GET("/resource", func(c *gin.Context) {
			c.String(http.StatusOK, "login")
		})
	}

	r.Run(":8080")
}
