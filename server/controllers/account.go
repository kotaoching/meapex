package controllers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/meapex/meapex/server/models"
	"github.com/meapex/meapex/server/utils"
)

func Signup(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := models.CreateUser(username, email, password)
	if err == nil {
		session.Set("userid", user.ID)
		session.Save()

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
				"status": "400",
				"source": "",
				"title":  "Invalid input",
				"detail": "Failed to register a new account",
			}},
		})
	}
}

func Signin(c *gin.Context) {
	session := sessions.Default(c)
	account := c.PostForm("account")
	password := c.PostForm("password")

	user, _ := models.GetUserByUsernameOrEmail(account)
	err := models.CheckPassword(user.Password, password)
	if err == nil {
		session.Set("userid", user.ID)
		session.Save()

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
				"status": "400",
				"source": "",
				"title":  "Failed to signin",
				"detail": "Invalid account or password.",
			}},
		})
	}
}

func Signout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userid")
	session.Save()
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
				"status": "200",
				"title":  "Find password success",
				"detail": "We have sent you an email, check your inbox.",
			},
		})
	} else {
		c.JSON(400, gin.H{
			"errors": []interface{}{map[string]interface{}{
				"status": "400",
				"source": "",
				"title":  "Failed to find password",
				"detail": "Invalid username or email.",
			}},
		})
	}
}
