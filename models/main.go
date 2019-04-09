package models

import (
	"database/sql"
)

//Genus represents a single genus
type Genus struct {
	GenusID int
	Genus   string
}

//Species represents a single species
type Species struct {
	SpeciesID  int            `db:"species_id"`
	Species    string         `db:"species"`
	CommonName sql.NullString `db:"common_name"`
	GenusID    int            `db:"genus_id"`
}

//Plants represents a particular plant in my collection
type Plants struct {
	PlantID   int
	SpeciesID sql.NullInt64
	DateAdded sql.NullFloat64
	Birthdate sql.NullFloat64
}

//SpeciesFlattened is like Species, but gives direct access to genus+species
type SpeciesFlattened struct {
	Species    string         `db:"species_name"`
	CommonName sql.NullString `db:"common_name"`
	SpeciesID  int            `db:"species_id"`
}
