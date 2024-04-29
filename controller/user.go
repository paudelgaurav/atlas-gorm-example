package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-api-permissions/constants"
	"github.com/paudelgaurav/gin-api-permissions/database"
	"github.com/paudelgaurav/gin-api-permissions/models"
	"github.com/paudelgaurav/gin-api-permissions/utils"
)

/*
Sample post data:

	{
		"username": sample username,
		"password": sample password,
		"permissions":"can_create_users, can_update_users, can_ping",
	}
*/
func CreateUser(c *gin.Context) {
	var (
		user        models.User
		permissions []string
	)
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(user.Permissions) > 0 {
		permissions = strings.Split(user.Permissions[0], ",")
		// removing unnecessary white spaces
		for i := range permissions {
			permissions[i] = strings.TrimSpace(permissions[i])
		}
	}

	// removing duplicate permissions
	permissions = utils.RemoveDuplicateStr(permissions)

	// checking if whether permissons are valid or not
	validPermissions := constants.GetAllPermissions()
	if !utils.SubSet(permissions, validPermissions) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid permission found"})
		return
	}

	user.Permissions = permissions
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "created"})
}

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
