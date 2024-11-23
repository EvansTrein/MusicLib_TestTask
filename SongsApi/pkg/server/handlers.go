package server

import (
	"SongsLib/SongsApi/pkg/database"
	myLog "SongsLib/SongsApi/pkg/logging"
	"SongsLib/SongsApi/pkg/models"
	"SongsLib/SongsApi/pkg/utils"

	"github.com/gin-gonic/gin"
)

func SongsHandler(ctx *gin.Context) {
	var songs []models.Song                   // переменная для возврата данных
	var params = make(map[string]string)      // сюда соберем параметры из запроса для их дальнейшей проверки
	databaseQuery := map[string]interface{}{} // сюда будем собирать фильтры для запроса в БД

	// получаем url запроса и оставляем только его путь
	urlStr := ctx.Request.URL.String()
	myLog.LogInfo.Println("Совершен запрос:", urlStr)

	// если запрос вообще без параметров, то возвращаем все данные и выходим из функции
	if urlStr == "/songs" {
		database.DB.Find(&songs)
		myLog.LogInfo.Println("Запрошены все данные")
		ctx.JSON(200, gin.H{"allData": songs})
		return
	}

	// получаем параметры из запроса
	ctx.BindQuery(&params)
	myLog.LogInfo.Println("Параметры запроса:", params)

	// проверяем параметры для фильтрации
	if value, ok := params["group"]; ok {
		databaseQuery["music_group"] = value
	}

	if value, ok := params["song"]; ok {
		databaseQuery["song_name"] = value
	}

	if value, ok := params["releaseDate"]; ok {
		databaseQuery["release_date"] = value
	}

	if value, ok := params["text"]; ok {
		databaseQuery["text"] = value
	}

	if value, ok := params["link"]; ok {
		databaseQuery["link"] = value
	}

	myLog.LogInfo.Println("Параметры после проверки:", databaseQuery)
	// проверяем, что данные могут применяться для фильтрации (т.е. в запрсое были поля нашей БД)
	if len(databaseQuery) == 0 {
		myLog.LogErr.Println("Переданные параметры запроса невалидны")
		ctx.JSON(400, gin.H{"error": "The passed request parameters are invalid"})
		return
	}

	// создаем запрос в БД на основе фильтров
	dbQuery := database.DB.Where(databaseQuery)

	// получаем параметры пагинации из запроса и проверяем их
	if value, ok := params["offset"]; ok {
		offset, err := utils.CheckOffset(value)
		if err != nil {
			myLog.LogErr.Println("В параметр offset пришел не integer")
			ctx.JSON(400, gin.H{"error": "Invalid query parameters, 'offset' must be positive integer."})
			return
		} else {
			dbQuery.Offset(offset) // добавляем к запросу offset если он есть и прошел проверку
		}
	}

	if value, ok := params["limit"]; ok {
		limit, err := utils.CheckLimit(value)
		if err != nil {
			myLog.LogErr.Println("В параметр limit пришел не integer")
			ctx.JSON(400, gin.H{"error": "Invalid query parameters, 'limit' must be positive integer."})
			return
		} else {
			dbQuery.Limit(limit) // добавляем к запросу limit если он есть и прошел проверку
		}
	}

	// выполяем запрос к БД
	if err := dbQuery.Find(&songs).Error; err != nil {
		myLog.LogErr.Println("Ошибка при выполнении запроса к базе данных:", err)
		ctx.JSON(500, gin.H{"error": "Error when executing a database query"})
		return
	}

	myLog.LogInfo.Println("Данные успешно отправлены")
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
