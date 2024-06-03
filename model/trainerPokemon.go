package model

import (
	"github.com/jinzhu/gorm"
)

type TrainerPokemon struct {
	gorm.Model
	UserID    string `json:"user_id"`
	PokemonID int    `json:"pokemon_id"`
	Level     int    `json:"level"`
}
