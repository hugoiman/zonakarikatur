package models

import "zonakarikatur/db"

// About is class
type About struct {
	IDAbout     int    `json:"idAbout"`
	Description string `json:"description"`
}

// GetAbout is func
func GetAbout() About {
	con := db.Connect()
	query := "SELECT idAbout, description FROM about"

	about := About{}

	_ = con.QueryRow(query).Scan(
		&about.IDAbout, &about.Description)

	defer con.Close()
	return about
}

// UpdateAbout is func
func UpdateAbout(idAbout string, data About) error {
	con := db.Connect()
	query := "UPDATE about SET description = ?"
	_, err := con.Exec(query, data.Description)

	defer con.Close()

	return err
}
