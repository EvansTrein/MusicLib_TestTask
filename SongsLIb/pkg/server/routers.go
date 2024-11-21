package server

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"

)

func InitRotes() {
	router := gin.Default()

	// Получение данных библиотеки с фильтрацией по всем полям и пагинацией
	router.GET("/songs", SongHandler)

	// Получение текста песни с пагинацией по куплетам
	router.GET("/song/:id/couplets", SongCoupletsHandler)

	// Добавление новой песни в формате JSON через сторонние API
	router.POST("/song", CreateSongHandler)

	// Изменение данных песни
	router.PUT("/song/:id/update", UpdateSongHandler)

	// Удаление песни
	router.DELETE("/song/:id/delete", DeleteSongHandler)

	// Сгенерировать сваггер на реализованное АПИ
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":3000")
}
