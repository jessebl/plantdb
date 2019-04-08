package main

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/jessebl/plantdb/models"
)

func speciesFlattened(db *sqlx.DB) ([]models.SpeciesFlattened, error) {
	var sS []models.SpeciesFlattened
	err := db.Select(&sS, "SELECT * FROM species_flattened;")
	return sS, err
}

func species(db *sqlx.DB) ([]models.Species, error) {
	var sS []models.Species
	err := db.Select(&sS, "SELECT * FROM species;")
	return sS, err
}

func setSpeciesProperty(db *sqlx.DB, speciesId int, property string, value string) (int64, error) {
	switch property {
	case "SpeciesId":
		property = "species_id"
	case "Species":
		property = "species"
	case "CommonName":
		property = "common_name"
	case "GenusId":
		property = "genus_id"
	default:
		return 0, errors.New("Invalid property")
	}
	stmt := "UPDATE species SET " + property + " = ? WHERE species_id = ?;"
	res, err := db.Exec(stmt, value, speciesId)
	if err != nil {
		return 0, err
	}
	rA, _ := res.RowsAffected()
	return rA, err
}

func main() {
	db, _ := sqlx.Connect("sqlite3", "./plant.db")
	defer db.Close()
	sS, _ := speciesFlattened(db)
	for _, s := range sS {
		fmt.Println(s)
	}
	res, err := setSpeciesProperty(db, 2, "SpeciesId", "5")
	fmt.Println("Rows affected:", res, "Errors:", err)
	sS, _ = speciesFlattened(db)
	for _, s := range sS {
		fmt.Println(s)
	}
}
