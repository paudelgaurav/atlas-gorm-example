package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/atlas-gorm-example/database"
	"github.com/paudelgaurav/atlas-gorm-example/models"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Basic Authentication credentials from the request
		username, password, hasAuth := c.Request.BasicAuth()
		if !hasAuth {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Fetch the user from the database
		var user models.User
		if err := database.DB.First(&user, "username = ?", username).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the provided password matches the user's password
		if !user.CheckPassword(password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// If all checks pass, set the user ID in the context for future use
		c.Set("userID", user.ID)
		c.Next()
	}
}
