package server

import (
	"SongsLib/SongsApi/pkg/database"
	"log"
)

func InitServer() {
	// Database initialization
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("Database connection error: ", errDatabase)
	} else {
		log.Println("Successful connection to the database")
		// тут сделать миграцию БД
	}
}

func StartServer() {
	InitRotes()
}
