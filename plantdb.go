package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/jessebl/plantdb/models"
)

func getSpeciesFlattened(db *sqlx.DB) []models.SpeciesFlattened {
	rows, err := db.Queryx("SELECT * FROM species_flattened")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer rows.Close()
	sS := []models.SpeciesFlattened{}
	for rows.Next() {
		s := models.SpeciesFlattened{}
		err = rows.StructScan(&s)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		sS = append(sS, s)
	}
	return sS
}

func getSpecies(db *sqlx.DB) []models.Species {
	rows, err := db.Queryx("SELECT * FROM species;")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer rows.Close()
	var sS []models.Species
	for rows.Next() {
		s := models.Species{}
		_ = rows.StructScan(&s)
		sS = append(sS, s)
	}
	return sS
}

func main() {
	db, err := sqlx.Connect("sqlite3", "./plant.db")
	defer db.Close()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	sS := getSpeciesFlattened(db)
	for _, s := range sS {
		fmt.Println(s)
	}
}
