package server

import (
	"SongsLib/SongsLIb/cmd/database"
	"log"
)

func InitServer() {
	// Database initialization
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("Database connection error: ", errDatabase)
	} else {
		log.Println("Successful connection to the database")
	}
}

func StartServer() {
	InitRotes()
}
