package database

import (
	"log"
	"pokemonApi/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance
var TrainerDatabase DbInstance
var TrainerPokemonDatabase DbInstance

func ConnectDb() {
	db, err := gorm.Open("sqlite3", "pokemons.db")
	if err != nil {
		panic(err)
	}

	log.Println("Database connection successful")
	db.AutoMigrate(&model.Pokemon{})

	Database = DbInstance{Db: db}
}

func ConnectTrainerDb() {
	db, err := gorm.Open("sqlite3", "trainer.db")
	if err != nil {
		panic(err)
	}

	log.Println("Trainer Database connection successful")
	db.AutoMigrate(&model.Trainer{})

	TrainerDatabase = DbInstance{Db: db}
}

func ConnectTrainerPokemonDb() {
	db, err := gorm.Open("sqlite3", "trainerPokemon.db")
	if err != nil {
		panic(err)
	}

	log.Println("Trainer pokemon Database connection successful")
	db.AutoMigrate(&model.TrainerPokemon{})

	TrainerPokemonDatabase = DbInstance{Db: db}
}
