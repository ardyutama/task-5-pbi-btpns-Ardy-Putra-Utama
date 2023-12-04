package database

import (
	"rakamin-golang/helpers"
	"rakamin-golang/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	helpers.LoadEnvVariables()
	dbUser := helpers.GetEnvVariable("DB_USER")
	dbPassword := helpers.GetEnvVariable("DB_PASSWORD")
	dbHost := helpers.GetEnvVariable("DB_HOST")
	dbPort := helpers.GetEnvVariable("DB_PORT")
	dbName := helpers.GetEnvVariable("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	DB = db

	db.AutoMigrate(&models.User{}, &models.Photo{})
}
