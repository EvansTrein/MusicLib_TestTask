package server

import (
	"SongsLib/SongsApi/pkg/models"
	myLog "SongsLib/SongsApi/pkg/logging"

	"github.com/gin-gonic/gin"
)

func SongHandler(ctx *gin.Context) {

}

func SongCoupletsHandler(ctx *gin.Context) {

}

func CreateSongHandler(ctx *gin.Context) {
	// var song models.Song  // переменная для хранения структуры, которую позжу будем записывать в postgres
	var req models.RequestData // переменная для запроса пришедшего на севрвер

	// парсим данные из тела запроса
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		myLog.LogErr.Println("ERROR: Invalid request body")
		ctx.JSON(404, gin.H{"error": "Incorrect data in body"})
		return
	} else {
		myLog.LogInfo.Printf("Group: %s, Song: %s", req.Group, req.Song)
	}

}

func UpdateSongHandler(ctx *gin.Context) {

}

func DeleteSongHandler(ctx *gin.Context) {

}
