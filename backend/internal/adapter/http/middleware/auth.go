package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("auth_token")

		if err != nil {
			log.Println(token, err.Error())
		}

		if err != nil || len(token) <= 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
