package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/meapex/meapex/server/db"
	"github.com/meapex/meapex/server/models"
	"github.com/meapex/meapex/server/package/token"
	"github.com/meapex/meapex/server/utils"
	"time"
)

func Signup(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := models.CreateUser(username, email, password)
	if err == nil {
		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"id": user.ID,
				"attributes": map[string]interface{}{
					"username":   user.Username,
					"email":      user.Email,
					"created_at": user.CreatedAt,
					"updated_at": user.UpdatedAt,
				},
			},
		})
	} else {
		c.JSON(400, gin.H{
			"errors": []interface{}{map[string]interface{}{
				"status":  "400",
				"title":   "Invalid input",
				"message": "Failed to register a new account",
			}},
		})
	}
}

func Signin(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")

	user, _ := models.GetUserByUsernameOrEmail(account)
	err := models.CheckPassword(user.Password, password)
	if err == nil {
		signedToken := token.New(user.ID, user.Username, time.Now().Add(time.Hour*24*7).Unix())

		redisConn := db.RedisPool.Get()
		redisConn.Do("SET", "me:user:"+user.Username, signedToken, "EX", int64(time.Hour*24*7))

		defer redisConn.Close()

		c.JSON(200, gin.H{
			"token": signedToken,
			"data": map[string]interface{}{
				"id":         user.ID,
				"username":   user.Username,
				"email":      user.Email,
				"created_at": user.CreatedAt,
				"updated_at": user.UpdatedAt,
			},
		})
	} else {
		c.JSON(400, gin.H{
			"errors": []interface{}{map[string]interface{}{
				"status":  "400",
				"title":   "Failed to signin",
				"message": "Invalid account or password.",
			}},
		})
	}
}

func Signout(c *gin.Context) {
	claims, err := token.Parse(c.Request)
	if err == nil {
		username := claims["username"].(string)

		redisConn := db.RedisPool.Get()
		redisConn.Do("DEL", "me:user:"+username)

		defer redisConn.Close()
	}
}

func FindPassword(c *gin.Context) {
	account := c.PostForm("account")

	user, err := models.GetUserByUsernameOrEmail(account)
	if err == nil {
		token := utils.GenerateRandom(16)
		user.Token = token
		models.UpdateUser(user)

		c.JSON(200, gin.H{
			"data": map[string]interface{}{
				"status":  "200",
				"title":   "Find password success",
				"message": "We have sent you an email, check your inbox.",
			},
		})
	} else {
		c.JSON(400, gin.H{
			"errors": []interface{}{map[string]interface{}{
				"status":  "400",
				"title":   "Failed to find password",
				"message": "Invalid username or email.",
			}},
		})
	}
}
