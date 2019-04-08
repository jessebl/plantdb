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

//structFieldToColumnName gets db column name from `db` tag on the given struct
//NOTE: strct must be a pointer to a struct
//TODO: Enforce the above limitation
func structFieldToColumnName(strct interface{}, field string) (string, error) {
	var dbProperty string
	fields, ok := reflect.TypeOf(strct).Elem().FieldByName(field)
	if !ok {
		return "", errors.New("Invalid property")
	}
	dbProperty, ok = fields.Tag.Lookup("db")
	if !ok {
		return "", errors.New("Unable to get value from `db` struct tag for field '" + field + "'")
	}
	return dbProperty, nil

}

func setSpeciesProperty(db *sqlx.DB, speciesId int, property string, value string) (int64, error) {
	dummy := models.Species{}
	colName, err := structFieldToColumnName(&dummy, property)
	stmt := "UPDATE species SET " + colName + " = ? WHERE species_id = ?;"
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
