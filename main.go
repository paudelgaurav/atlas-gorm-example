package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/paudelgaurav/gin-api-permissions/controller"
	"github.com/paudelgaurav/gin-api-permissions/database"
	"github.com/paudelgaurav/gin-api-permissions/middleware"
	"github.com/paudelgaurav/gin-api-permissions/models"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	log.Print("User Table Migrate")
	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("problem while migrating user table")
	}
}

func serveApplication() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.POST("/user", middleware.BasicAuthPermission("can_create_users"), controller.CreateUser)
	r.GET("/users", middleware.BasicAuthPermission("can_list_users"), controller.GetUsers)

	r.Run()
}
