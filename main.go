package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/paudelgaurav/atlas-gorm-example/controller"
	"github.com/paudelgaurav/atlas-gorm-example/database"
	"github.com/paudelgaurav/atlas-gorm-example/models"
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
	r.POST("/users", controller.CreateUser)
	r.GET("/users", controller.GetUsers)

	r.Run()
}
