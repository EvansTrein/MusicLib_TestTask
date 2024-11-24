package server

import (
	"SongsLib/SongsApi/pkg/database"
	myLog "SongsLib/SongsApi/pkg/logging"
	"SongsLib/SongsApi/pkg/models"
)

func InitServer() {
	// Подключение к БД
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		myLog.LogErr.Fatal("Database connection error: ", errDatabase)
	} else {
		myLog.LogInfo.Println("Successful connection to the database")
		// структура БД должна быть создана путем миграций при старте сервиса
		database.DB.AutoMigrate(&models.Song{})
	}
}

func StartServer() {
	InitRotes()
}
