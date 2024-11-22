package server

import (
	"SongsLib/SongsApi/pkg/database"
	myLog "SongsLib/SongsApi/pkg/logging"
	"SongsLib/SongsApi/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SongsHandler(ctx *gin.Context) {
	var songs []models.Song
	databaseQuery := map[string]interface{}{} // сюда собирать фильтры для запроса в БД

	urlStr := ctx.Request.URL.String()
	myLog.LogInfo.Println("Совершен запрос:", urlStr)

	// получаем параметры фильтрации из запроса
	group := ctx.Query("group")
	songName := ctx.Query("songName")
	releaseDate := ctx.Query("releaseDate")
	text := ctx.Query("text")
	link := ctx.Query("link")

	// получаем параметры пагинации из запроса
	offsetStr := ctx.Query("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		myLog.LogErr.Println("В параметр offset пришел не integer")
		ctx.JSON(400, gin.H{"error": "Invalid query parameters, 'offset' must be positive integer."})
		return
	}
	
	limitStr := ctx.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		myLog.LogErr.Println("В параметр limit пришел не integer")
		ctx.JSON(400, gin.H{"error": "Invalid query parameters, 'limit' must be positive integer."})
		return
	}

	myLog.LogInfo.Println(group)
	myLog.LogInfo.Println(songName)
	myLog.LogInfo.Println(releaseDate)
	myLog.LogInfo.Println(text)
	myLog.LogInfo.Println(link)
	myLog.LogInfo.Println(limit)
	myLog.LogInfo.Println(offset)

	databaseQuery["music_group"] = group
	// databaseQuery["song_name"] = songName
	database.DB.Where(databaseQuery).Find(&songs)

	// database.DB.Where("music_group = ?", group).Offset(offset).Limit(limit).Find(&songs)

	for _, el := range songs {
		myLog.LogInfo.Println(el.MusicGroup, el.SongName)
	}

	ctx.JSON(200, gin.H{"data": songs})
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
		myLog.LogInfo.Printf("на сервер пришли данные - Group: %s, Song: %s", req.MusicGroup, req.Song)
	}

	// сохраняем данные в поля таблицы
	songDb.MusicGroup = req.MusicGroup
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
