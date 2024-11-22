package server

import (
	"SongsLib/SongsApi/pkg/database"
	myLog "SongsLib/SongsApi/pkg/logging"
	"SongsLib/SongsApi/pkg/models"

	"github.com/gin-gonic/gin"
)

func SongHandler(ctx *gin.Context) {

}

func SongCoupletsHandler(ctx *gin.Context) {

}

func CreateSongHandler(ctx *gin.Context) {
	var songDb models.Song     // переменная для хранения структуры, которую позже будем записывать в postgres
	var req models.RequestData // переменная для запроса пришедшего на севрвер
	// var songData models.SongData

	// парсим данные из тела запроса
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		myLog.LogErr.Println("ERROR: Invalid request body")
		ctx.JSON(404, gin.H{"error": "Incorrect data in body"})
		return
	} else {
		myLog.LogInfo.Printf("на сервер пришли данные - Group: %s, Song: %s", req.Group, req.Song)
	}

	// сохраняем данные в поля таблицы 
	songDb.Group = req.Group
	songDb.SongName = req.Song
	songDb.ReleaseDate = "16.07.2006"
	songDb.Text = `Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight`
	songDb.Link = "https://www.youtube.com/watch?v=Xsp3_a-PMTw3"

	// создаем запись в таблице
	entryDb := database.DB.Create(&songDb)
	if entryDb.Error != nil {
		myLog.LogErr.Println("Ошибка при записи в базу данных", entryDb.Error)
		ctx.JSON(500, gin.H{"error": "failed to save to the database"})
		return
	} else {
		myLog.LogInfo.Println("Песня сохранена успешно")
		ctx.JSON(201, gin.H{"message": "The song has been successfully created"})
	}
}

func UpdateSongHandler(ctx *gin.Context) {

}

func DeleteSongHandler(ctx *gin.Context) {

}
