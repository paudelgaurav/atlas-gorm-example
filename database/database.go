package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// NewDatabase creates a new database instance
func Connect() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port)

	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		panic(err)
	}

	if err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error; err != nil {
		panic(err)
	}

	// close the current connection
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	if dbErr := sqlDb.Close(); dbErr != nil {
		panic(dbErr)
	}

	// reopen connection with the given database, after creating or checking if the database exists

	urlWithDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err = gorm.Open(mysql.Open(urlWithDB))
	if err != nil {
		panic(err)
	}

	conn, err := db.DB()
	if err != nil {
		panic(err)
	}

	conn.SetConnMaxLifetime(time.Minute * 5)
	conn.SetMaxOpenConns(5)
	conn.SetMaxIdleConns(1)

	DB = db
}
