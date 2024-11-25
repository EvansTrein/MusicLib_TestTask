package server

import (
	"SongsLib/SongsApi/pkg/database"
	"SongsLib/SongsApi/pkg/envs"
	myLog "SongsLib/SongsApi/pkg/logging"
	"SongsLib/SongsApi/pkg/models"
)

func InitServer() {
	// Инициализация внешних значений ENV
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		myLog.LogErr.Fatal("Ошибка загрузки ENV: ", errEnvs)
	} else {
		myLog.LogInfo.Println("Успешное получение ENV")
	}
	// Подключение к БД
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		myLog.LogErr.Fatal("Ошибка подключения к базе данных: ", errDatabase)
	} else {
		myLog.LogInfo.Println("Успешное подключение к базе данных")
		// структура БД должна быть создана путем миграций при старте сервиса
		database.DB.AutoMigrate(&models.Song{})
	}
}

func StartServer() {
	InitRotes()
}
