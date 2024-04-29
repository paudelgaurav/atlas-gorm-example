package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/paudelgaurav/atlas-gorm-example/controller"
	"github.com/paudelgaurav/atlas-gorm-example/database"
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

}

func serveApplication() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.POST("/users", controller.CreateUser)
	r.GET("/users", controller.GetUsers)

	r.Run()
}
