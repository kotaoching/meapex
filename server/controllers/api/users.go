package api

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/meapex/meapex/server/models"
)

func Me(c *gin.Context) {
	session := sessions.Default(c)
	userid := session.Get("userid")
	if userid != nil {
		user, err := models.GetUserById(userid)

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
			c.JSON(401, gin.H{})
		}
	} else {
		c.JSON(401, gin.H{
			"errors": []interface{}{map[string]interface{}{
				"status": "401",
				"source": "",
				"title":  "Unauthorized",
				"detail": "Authorization is required.",
			}},
		})
	}
}
