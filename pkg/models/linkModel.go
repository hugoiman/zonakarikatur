package models

import "zonakarikatur/db"

// Link is class
type Link struct {
	IDLink int    `json:"idLink"`
	Link   string `json:"link"`
}

// GetLink is func
func GetLink() Link {
	con := db.Connect()
	query := "SELECT idLink, link FROM link"

	link := Link{}

	_ = con.QueryRow(query).Scan(
		&link.IDLink, &link.Link)

	defer con.Close()
	return link
}

// UpdateLink is func
func UpdateLink(idLink string, data Link) error {
	con := db.Connect()
	query := "UPDATE link SET link = ?"
	_, err := con.Exec(query, data.Link)

	defer con.Close()

	return err
}
