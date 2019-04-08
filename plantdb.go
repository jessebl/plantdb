package main

import (
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

func setCommonName(db *sqlx.DB, speciesId int, cN string) (int64, error) {
	q := `UPDATE species SET common_name = ? WHERE species_id = ?`
	res, err := db.Exec(q, cN, speciesId)
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
	res, err := setCommonName(db, 2, "yee")
	fmt.Println("Rows affected:", res, "Errors:", err)
	sS, _ = speciesFlattened(db)
	for _, s := range sS {
		fmt.Println(s)
	}
}
