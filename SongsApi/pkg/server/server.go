package server

import (
	"SongsLib/SongsApi/pkg/database"
	myLog "SongsLib/SongsApi/pkg/logging"
	"SongsLib/SongsApi/pkg/models"
)

func InitServer() {
	// Database initialization
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		myLog.LogErr.Fatal("Database connection error: ", errDatabase)
	} else {
		myLog.LogInfo.Println("Successful connection to the database")
		// структура БД должна быть создана путем миграций при старте сервиса
		database.DB.AutoMigrate(&models.Song{})
		database.DB.Exec("ALTER TABLE songs RENAME COLUMN group TO MusicGroup")
	}
}

func StartServer() {
	InitRotes()
}
