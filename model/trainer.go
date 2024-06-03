package model

import (
	"github.com/jinzhu/gorm"
)

type Trainer struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password string `json:"password"`
	Level    int    `json:"level"`
	Coins    int    `json:"coins"`
}
