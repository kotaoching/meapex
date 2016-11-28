package server

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/meapex/meapex/server/controllers"
	"github.com/meapex/meapex/server/controllers/api"
	"net/http"
)

func Init() {
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	//store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("abc_sess", store))

	r.POST("/signup", controllers.Signup)
	r.POST("/signin", controllers.Signin)
	r.POST("/signout", controllers.Signout)
	r.POST("/find-password", controllers.FindPassword)

	apiv1 := r.Group("/api")
	{
		apiv1.GET("/me", api.Me)
		apiv1.GET("/resource", func(c *gin.Context) {
			c.String(http.StatusOK, "login")
		})
	}

	r.Run(":8080")
}
