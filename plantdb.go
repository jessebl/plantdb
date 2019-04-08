package main

import (
	"errors"
	"fmt"
	"reflect"

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
	//Get db name for property from `db` tag on Species struct
	var dbProperty string
	dummy := models.Species{}
	fields, ok := reflect.TypeOf(&dummy).Elem().FieldByName(property)
	if !ok {
		return 0, errors.New("Invalid property")
	}
	dbProperty, ok = fields.Tag.Lookup("db")
	if !ok {
		return 0, errors.New("Unable to get value from `db` struct tag for field '" + property + "'")
	}
	stmt := "UPDATE species SET " + dbProperty + " = ? WHERE species_id = ?;"
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
	res, err := setSpeciesProperty(db, 2, "CommonName", "Super duper old man of the Andes")
	fmt.Println("Rows affected:", res, "Errors:", err)
	sS, _ = speciesFlattened(db)
	for _, s := range sS {
		fmt.Println(s)
	}
}
