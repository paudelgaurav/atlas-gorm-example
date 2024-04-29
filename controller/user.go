package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/atlas-gorm-example/database"
	"github.com/paudelgaurav/atlas-gorm-example/models"
	"gorm.io/gorm"
)

/*
Sample post data:

	{
		"username": sample username,
	}
*/
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err

		}
		return tx.Create(&models.UserRole{UserID: user.ID, Role: "general"}).Error
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "created"})
}

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := database.DB.Preload("Role").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
