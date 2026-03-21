package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title 		string `json:"judul"`
	Content		string `json:"isi"`
	Category 	string `json:"kategori"`

}