package models

import (
	"database/sql"
)

type Genus struct {
	GenusId int
	Genus   string
}

type Species struct {
	SpeciesId  int            `db:"species_id"`
	Species    string         `db:"species"`
	CommonName sql.NullString `db:"common_name"`
	GenusId    int            `db:"genus_id"`
}

type Plants struct {
	PlantId   int
	SpeciesId sql.NullInt64
	DateAdded sql.NullFloat64
	Birthdate sql.NullFloat64
}

type SpeciesFlattened struct {
	Species    string         `db:"species_name"`
	CommonName sql.NullString `db:"common_name"`
	SpeciesId  int            `db:"species_id"`
}
