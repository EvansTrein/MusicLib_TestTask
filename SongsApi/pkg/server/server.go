package server

import (
	"SongsLib/SongsApi/pkg/database"
	myLog "SongsLib/SongsApi/pkg/logging"
)

func InitServer() {
	// Database initialization
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		myLog.LogErr.Fatal("Database connection error: ", errDatabase)
	} else {
		myLog.LogInfo.Println("Successful connection to the database")
		// тут сделать миграцию БД
	}
}

func StartServer() {
	InitRotes()
}
