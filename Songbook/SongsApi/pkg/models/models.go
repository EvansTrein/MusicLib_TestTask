package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	MusicGroup  string `gorm:"not null"`
	SongName    string `gorm:"not null"`
	ReleaseDate string
	Text        string
	Link        string
}

type SongData struct {
	MusicGroup  string `json:"group" example:"Muse"`
	SongName    string `json:"song" example:"Supermassive Black Hole"`
	ReleaseDate string `json:"releaseDate" example:"16.07.2006"`
	Text        string `json:"text" example:"\"Ooh baby, don't you know I suffer?\\nOoh baby, can you hear me moan?\\nYou caught me under false pretenses\\nHow long before you let me go?\\n\\nOoh\\nYou set my soul alight\\nOoh\\nYou set my soul alight\""`
	Link        string `json:"link" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"`
}

type RequestData struct {
	MusicGroup string `json:"group" binding:"required" example:"Muse"`
	Song       string `json:"song" binding:"required" example:"Supermassive Black Hole"`
}

type DataFromAPI struct {
	ReleaseDate string `json:"releaseDate" binding:"required" example:"16.07.2006"`
	Text        string `json:"text" binding:"required" example:"\"Ooh baby, don't you know I suffer?\\nOoh baby, can you hear me moan?\\nYou caught me under false pretenses\\nHow long before you let me go?\\n\\nOoh\\nYou set my soul alight\\nOoh\\nYou set my soul alight\""`
	Link        string `json:"link" binding:"required" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"`
}

type ErrResponce struct {
	ErrorMessage string
}

type ResponceData struct {
	Message string
}
