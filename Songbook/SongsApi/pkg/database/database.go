package database

import (
	"fmt"

	"SongsLib/SongsApi/pkg/envs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		envs.ServerEnvs.POSTGRES_HOST, envs.ServerEnvs.POSTGRES_USER, envs.ServerEnvs.POSTGRES_PASSWORD, envs.ServerEnvs.POSTGRES_NAME, 
		envs.ServerEnvs.POSTGRES_PORT, envs.ServerEnvs.POSTGRES_USE_SSL)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return err
	} else {
		DB = db
		return nil
	}
}
