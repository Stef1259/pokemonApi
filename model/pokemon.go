package model

import (
	"github.com/jinzhu/gorm"
)

type Pokemon struct {
	gorm.Model
	Name  string `json:"name"`
	Type  string `json:"type"`
	Moves string `json:"moves"`
}
