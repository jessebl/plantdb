package models

import "time"

type Genus struct {
	GenusId int
	Genus   string
}

type Species struct {
	SpeciesId  int
	Species    string
	CommonName string
	GenusId    int
}

type Plants struct {
	PlantId   int
	SpeciesId int
	DateAdded time.Time
	Birthdate time.Time
}

type SpeciesFlattened struct {
	Species, CommonName string
	SpeciesId           int
}
