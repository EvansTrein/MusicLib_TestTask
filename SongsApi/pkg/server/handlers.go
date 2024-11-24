package server

import (
	"SongsLib/SongsApi/pkg/database"
	myLog "SongsLib/SongsApi/pkg/logging"
	"SongsLib/SongsApi/pkg/models"
	"SongsLib/SongsApi/pkg/utils"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SongsHandler(ctx *gin.Context) {
	var songs []models.Song                      // переменная для возврата данных
	var params = make(map[string]string)         // сюда соберем параметры из запроса для их дальнейшей проверки
	databaseQuery := map[string]interface{}{}    // сюда будем собирать фильтры для запроса в БД
	dbQuery := database.DB.Model(&models.Song{}) // создаем запрос для БД, который будем наполнять параметрами
	var wg sync.WaitGroup                        // счетчик для контроля горутин

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

	// запускаем проверку парметров запроса в отдельной горутинее и идем далее по другим делам
	wg.Add(1)
	go func() {
		defer wg.Done()
		if value, ok := params["group"]; ok {
			databaseQuery["music_group"] = value
			delete(params, "group")
		}

		if value, ok := params["song"]; ok {
			databaseQuery["song_name"] = value
			delete(params, "song")
		}

		if value, ok := params["releaseDate"]; ok {
			databaseQuery["release_date"] = value
			delete(params, "releaseDate")
		}

		if value, ok := params["text"]; ok {
			databaseQuery["text"] = value
			delete(params, "text")
		}

		if value, ok := params["link"]; ok {
			databaseQuery["link"] = value
			delete(params, "link")
		}
	}()

	// получаем параметры пагинации из запроса и проверяем их
	if value, ok := params["offset"]; ok {
		delete(params, "offset")
		offset, err := utils.CheckOffset(value) // проверка значения offset и преобразование его к типу int (из запроса мы получили string)
		if err != nil {
			myLog.LogErr.Println("В параметр offset пришел не integer")
			ctx.JSON(400, gin.H{"error": "Invalid query parameters, 'offset' or 'limit', they must be positive integers."})
			wg.Wait() // если offset не прошел проверку, то ждем все горутины, прежде чем выйти из функции
			return
		} else {
			dbQuery.Offset(offset) // добавляем к запросу offset если он есть и прошел проверку
		}
	}

	if value, ok := params["limit"]; ok {
		delete(params, "limit")
		limit, err := utils.CheckLimit(value) // проверка значения limit и преобразование его к типу int (из запроса мы получили string)
		if err != nil {
			myLog.LogErr.Println("В параметр limit пришел не integer")
			ctx.JSON(400, gin.H{"error": "Invalid query parameters, 'offset' or 'limit', they must be positive integers"})
			wg.Wait() // если limit не прошел проверку, то ждем все горутины, прежде чем выйти из функции
			return
		} else {
			dbQuery.Limit(limit) // добавляем к запросу limit если он есть и прошел проверку
		}
	}

	wg.Wait() // ждем все горутины

	myLog.LogInfo.Println("Параметры после проверки:", databaseQuery)
	// проверяем, что данные могут применяться для фильтрации (т.е. в запрсое были ТОЛЬКО разрешенные поля нашей БД)
	if len(params) != 0 {
		// ранее, мы убирали параметры, которые прошли проверку, но если мы попали сюда, значит в параметрах пришло не то, что ожидалось
		// даже если в параметрах были и разрешенные и неразрешенные вместе,
		// всеравно - делать лишний запрос к БД с фильтрами, которые БД не ожидает, я не стал
		myLog.LogErr.Println("Переданные параметры запроса невалидны")
		ctx.JSON(400, gin.H{"error": "The passed request parameters are invalid"})
		return
	} else if len(databaseQuery) > 0 {
		dbQuery.Where(databaseQuery) // добавляем фильтры к запросу в БД, если они есть
	}

	// выполяем запрос к БД с отражением в консоли SQL запроса
	result := dbQuery.Debug().Find(&songs)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	} else if len(songs) == 0 {
		ctx.JSON(404, gin.H{"error": "Record not found"})
		return
	} else {
		myLog.LogInfo.Println("Данные успешно отправлены")
		ctx.JSON(200, gin.H{"data": songs})
	}
}

func SongCoupletsHandler(ctx *gin.Context) {
	var song models.Song  // структура для песни которую будем возвращать
	id := ctx.Param("id") // получаем id из url
	offsetStr := ctx.Query("offset") // получаем начальный параметр из запроса
	limitStr := ctx.Query("limit") // получаем конечный параметр из запроса

	urlStr := ctx.Request.URL.String()
	myLog.LogInfo.Println("Совершен запрос:", urlStr)

	// выполяем запрос к БД для поиска нужной песни по id
	if result := database.DB.Debug().Where("id = ?", id).First(&song); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{"error": "User not found"})
		} else {
			ctx.JSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	myLog.LogInfo.Println("Найденная песня:", song)

	// получаем текст песни и разбиваем его по абзацам на слайс, 1 элемент - 1 абзац
	data := strings.Split(song.Text, `\n\n`)

	// если праметров для пагинации не передали, то возвращаем весь текст песни и выходим из функции
	if offsetStr == "" && limitStr == "" {
		ctx.JSON(200, gin.H{"textSong": data})
		return
	}

	// создаем мапу для хранения значений парметров фильтрации
	filterTextSong := make(map[string]int, 2)

	// проверка значения offset и преобразование его к типу int (из запроса мы получили string)
	if offsetStr != "" {
		offset, err := utils.CheckOffset(offsetStr)
		if err != nil || offset <= 0 {
			myLog.LogErr.Println("В параметр offset пришел не positive integers.")
			ctx.JSON(400, gin.H{"error": "Invalid query parameters, 'offset' or 'limit', they must be positive integers"})
			return
		} else {
			filterTextSong["startIndx"] = offset - 1 // для удобства (индексация начинается с 0)
		}
	}

	// проверка значения limit и преобразование его к типу int (из запроса мы получили string)
	if limitStr != "" {
		limit, err := utils.CheckLimit(limitStr)
		if err != nil || limit <= 0 {
			myLog.LogErr.Println("В параметр limit пришел не positive integers.")
			ctx.JSON(400, gin.H{"error": "Invalid query parameters, 'offset' or 'limit', they must be positive integers."})
			return
		} else {
			filterTextSong["endIndx"] = limit
		}
	}

	// проверяем какие параметры есть
	startValue, okStart := filterTextSong["startIndx"]
	endValue, okEnd := filterTextSong["endIndx"]

	// фильтрация текста песни по абзацам, с учетом индексов у слайса (нельзя обращаться к несуществующему индексу):
	// если запрошены start и end - дудет выведено от start по end ВКЛЮЧИТЕЛЬНО, есть проверка на логику
	// если запрошенен только start - дудет выведено от start и до конца
	// если запрошенен только end - дудет выведено от начала и до end ВКЛЮЧИТЕЛЬНО
	if okStart && okEnd {
		myLog.LogInfo.Println("Были запрошено 2 праметра, начало и конец")
		switch {
		case startValue + 1 > len(data):
			myLog.LogErr.Println("Были запрошены нивалидные номера куплетов")
			ctx.JSON(400, gin.H{"error": "This song has fewer verses"})
			return
		case startValue + 1 > endValue:
			myLog.LogErr.Println("Были запрошены нивалидные номера куплетов")
			ctx.JSON(400, gin.H{"error": "The starting verse number cannot be less than the ending verse number"})
			return
		case endValue > len(data) && startValue <= len(data):
			myLog.LogInfo.Println("Данные текста песни успешно отправлены")
			ctx.JSON(200, gin.H{"data": data[startValue:]})
			return
		case endValue <= len(data) && startValue <= endValue:
			myLog.LogInfo.Println("Данные текста песни успешно отправлены")
			ctx.JSON(200, gin.H{"data": data[startValue:endValue]})
			return
		}
	} else if okStart && !okEnd {
		myLog.LogInfo.Println("Было запрошено только с какого начинаем")
		switch {
		case startValue + 1 > len(data):
			myLog.LogErr.Println("Был запрошен нивалидный номер куплета")
			ctx.JSON(400, gin.H{"error": "This song has fewer verses"})
			return
		default:
			myLog.LogInfo.Println("Данные текста песни успешно отправлены")
			ctx.JSON(200, gin.H{"data": data[startValue:]})
			return
		}
	} else if !okStart && okEnd {
		myLog.LogInfo.Println("Было запрошено только до какого выводить")
		switch {
		case endValue > len(data):
			myLog.LogInfo.Println("Данные текста песни успешно отправлены")
			ctx.JSON(200, gin.H{"data": data})
			return
		default:
			myLog.LogInfo.Println("Данные текста песни успешно отправлены")
			ctx.JSON(200, gin.H{"data": data[:endValue]})
			return
		}
	} else {
		myLog.LogInfo.Println("Непредвиденное поведение, обратитесь к разработчику")
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
}

func CreateSongHandler(ctx *gin.Context) {
	var songDb models.Song     // переменная для хранения структуры, которую позже будем записывать в postgres
	var req models.RequestData // переменная для запроса пришедшего на севрвер
	// var songData models.SongData

	urlStr := ctx.Request.URL.String()
	myLog.LogInfo.Println("Совершен запрос:", urlStr)

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
	// songDb.Text = `Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight`
	songDb.Text = `Couplet - 1\n\nCouplet - 2\n\nCouplet - 3\n\nCouplet - 4\n\nCouplet - 5\n\nCouplet - 6\n\nCouplet - 7\n\nCouplet - 8`
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
	var song models.Song  // структура для песни которую будем удалять
	id := ctx.Param("id") // получаем id из url

	urlStr := ctx.Request.URL.String()
	myLog.LogInfo.Println("Совершен запрос:", urlStr)

	// выполяем запрос к БД для поиска нужной песни по id
	if result := database.DB.Where("id = ?", id).First(&song); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{"error": "User not found"})
		} else {
			ctx.JSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	// выполяем запрос к БД с отражением в консоли SQL запроса
	database.DB.Unscoped().Delete(&song)
	ctx.JSON(200, gin.H{"message": "user deleted successfully"})
}
