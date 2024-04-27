package model

import "gorm.io/gorm"

type Episode struct {
	gorm.Model
	Title string `gorm:"type:varchar(100)"`
}

type Config struct {
	Title           string
	Host            string
	ImgUrl          string
	Tagline         string
	Desc            string
	SpotifyUrl      string
	ApplePodcastUrl string
}
