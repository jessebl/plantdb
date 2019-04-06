PRAGMA foreign_keys = on;

CREATE TABLE genera (
	genus_id INTEGER PRIMARY KEY,
	genus TEXT
);

INSERT INTO genera (genus_id, genus) VALUES
	(1, "Mammillaria"),
	(2, "Sedum"),
	(3, "Codiaeum"),
	(4, "Euphorbia"),
	(5, "Oreocereus"),
	(6, "Gymnocalycium"),
	(7, "Cereus"),
	(8, "Agave"),
	(9, "Echeveria"),
	(10, "Graptosedum");

CREATE TABLE species (
	species_id INTEGER PRIMARY KEY,
	species TEXT,
	common_name TEXT,
	genus_id INTEGER NOT NULL,
	FOREIGN KEY(genus_id) REFERENCES genera(genus_id)
);

INSERT INTO species (genus_id, species) VALUES
	(1, "carmenae"),
	(2, "treleasei"),
	(3, "variegatum"),
	(4, "tirucalli"),
	(5, "trollii"),
	(6, "baldianum"),
	(7, "forbesii monstrose"),
	(8, "sisalana variegata"),
	(9, "'Perle von Nurnberg'"),
	(10, "'California Sunset'");

CREATE VIEW species_flattened AS
SELECT genus||" "||species,species_id
FROM species s
LEFT JOIN genera g ON s.genus_id = g.genus_id;
