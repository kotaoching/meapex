package api

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/meapex/meapex/server/db"
	"github.com/meapex/meapex/server/models"
	"github.com/meapex/meapex/server/package/token"
)

func Me(c *gin.Context) {
	claims, err := token.Parse(c.Request)
	if err == nil {
		userid := claims["userid"].(string)
		username := claims["username"].(string)

		redisConn := db.RedisPool.Get()
		exists, _ := redis.Bool(redisConn.Do("EXISTS", "me:user:"+username+":token"))
		if !exists {
			c.JSON(401, gin.H{
				"errors": []interface{}{map[string]interface{}{
					"status":  "401",
					"title":   "invalid_token",
					"message": "Invalid authorized token.",
				}},
			})
			c.Abort()
			return
		}

		defer redisConn.Close()

		user, err := models.GetUserById(userid)
		if err == nil {
			c.JSON(200, gin.H{
				"data": user,
			})
		} else {
			c.JSON(401, gin.H{})
		}
	} else {
		c.JSON(401, gin.H{
			"errors": []interface{}{map[string]interface{}{
				"status":  "401",
				"title":   "authentication_error",
				"message": "Authorization required.",
			}},
		})
	}
}
