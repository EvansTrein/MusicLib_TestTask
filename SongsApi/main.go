package main

import (
	"SongsLib/SongsApi/pkg/server"
)

// @title           Онлайн библиотека песен🎶
// @version         0.1
// @description     Тестовое задание от Effective Mobile

// @contact.name   Evans Trein
// @contact.email  evanstrein@icloud.com

// @host      localhost:3000

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func init() {
	server.InitServer()
}

func main() {
	server.InitRotes()
}
