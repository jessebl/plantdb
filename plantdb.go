package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type speciesFlattened struct {
	species, commonName string
}

func getSpeciesFlattened(db *sqlx.DB) []speciesFlattened {
	rows, err := db.Query("SELECT genus||' '||species, common_name FROM species s LEFT JOIN genera g ON s.genus_id = g.genus_id;")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer rows.Close()
	sS := []speciesFlattened{}
	for rows.Next() {
		var specialName, commonName string
		_ = rows.Scan(&specialName, &commonName)
		sS = append(sS, speciesFlattened{species: specialName,
			commonName: commonName})
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
