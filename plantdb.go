package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/jessebl/plantdb/models"
)

func getSpeciesFlattened(db *sqlx.DB) []models.SpeciesFlattened {
	rows, err := db.Query("SELECT genus||' '||species, species_id, common_name FROM species s LEFT JOIN genera g ON s.genus_id = g.genus_id;")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer rows.Close()
	sS := []models.SpeciesFlattened{}
	for rows.Next() {
		var specialName, commonName string
		var speciesId int
		_ = rows.Scan(&specialName, &speciesId, &commonName)
		sS = append(sS, models.SpeciesFlattened{Species: specialName,
			CommonName: commonName,
			SpeciesId:  speciesId})
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
