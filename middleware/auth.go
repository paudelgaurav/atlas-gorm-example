package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-api-permissions/database"
	"github.com/paudelgaurav/gin-api-permissions/models"
	"github.com/paudelgaurav/gin-api-permissions/utils"
)

func BasicAuthPermission(permission string) gin.HandlerFunc {
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
		if !utils.Exists(permission, user.Permissions) {
			c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			c.Abort()
			return
		}
		// If all checks pass, set the user ID in the context for future use
		c.Set("userID", user.ID)
		c.Next()
	}
}
