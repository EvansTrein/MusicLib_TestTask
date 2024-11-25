package main

import (
	"SongsLib/SongsApi/pkg/logging"
	"SongsLib/SongsApi/pkg/server"
)

// @title           Онлайн библиотека песен🎶
// @version         0.1
// @description     Тестовое задание от Effective Mobile

// @contact.name   Evans Trein
// @contact.email  evanstrein@icloud.com
// @contact.url  https://github.com/EvansTrein

// @host      localhost:3000
// @schemes   http

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func init() {
	logging.InitLogger()
	server.InitServer()
}

func main() {
	server.InitRotes()
}
