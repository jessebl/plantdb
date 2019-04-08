package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/jessebl/plantdb/models"
)

func getSpeciesFlattened(db *sqlx.DB) ([]models.SpeciesFlattened, error) {
	var sS []models.SpeciesFlattened
	err := db.Select(&sS, "SELECT * FROM species_flattened;")
	return sS, err
}

func getSpecies(db *sqlx.DB) ([]models.Species, error) {
	var sS []models.Species
	err := db.Select(&sS, "SELECT * FROM species;")
	return sS, err
}

func main() {
	db, err := sqlx.Connect("sqlite3", "./plant.db")
	defer db.Close()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	sS, _ := getSpeciesFlattened(db)
	for _, s := range sS {
		fmt.Println(s)
	}
}
