package db

import (
	"fmt"
	"fr33d0mz/moneyflowx/logger"
	"fr33d0mz/moneyflowx/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func initDB() *gorm.DB {
	settingsParam := utils.AppSettings.PostgresParams

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Dushanbe",
		settingsParam.Server, settingsParam.User, settingsParam.Password, settingsParam.Database,
		settingsParam.Port, settingsParam.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error.Fatal("Failed to connect to postgreSQL database")
	}

	return db
}

func StartDbConnection() {
	database = initDB()
}

func GetDBConn() *gorm.DB {
	return database
}

func DisconnectDB(db *gorm.DB) {
	_db, err := db.DB()
	if err != nil {
		logger.Error.Fatal("Failed to kill connection from database. Error is: ", err.Error())
	}

	_db.Close()
}
