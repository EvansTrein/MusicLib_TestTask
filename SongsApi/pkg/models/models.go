package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Group       string `gorm:"not null"`
	SongName    string `gorm:"not null"`
	ReleaseDate string
	Text        string
	Link        string
}

type RequestData struct {
	Group string `json:"group"`
    Song  string `json:"song"`
}