package models

import (
	"fmt"
	"time"

	"zonakarikatur/db"
)

// Galleries is func
type Galleries struct {
	Galleries []Gallery `json:"galleries"`
}

// Gallery is func
type Gallery struct {
	IDGallery int    `json:"idGallery"`
	Image     string `json:"image" validate:"required"`
	Category  string `json:"category" validate:"required,eq=karikatur-full-body|eq=karikatur-half-body|eq=sketsa|eq=mozaik|eq=ilustrasi-wajah|eq=vexel|eq=kolase"`
	CreatedAt string `json:"createdAt"`
}

// GetGalleries is func
func GetGalleries(limit, offset int) Galleries {
	fmt.Println(limit, offset)
	con := db.Connect()
	query := "SELECT idGallery, image, category, createdAt FROM gallery ORDER BY idGallery DESC LIMIT ?,?"
	rows, _ := con.Query(query, limit, offset)

	var createdAt time.Time
	gallery := Gallery{}
	galleries := Galleries{}

	for rows.Next() {
		rows.Scan(
			&gallery.IDGallery, &gallery.Image, &gallery.Category, &createdAt,
		)

		gallery.CreatedAt = createdAt.Format("02 Jan 2006")
		galleries.Galleries = append(galleries.Galleries, gallery)
	}

	defer con.Close()

	return galleries
}

// GetGallery is func
func GetGallery(idGallery string) (Gallery, error) {
	con := db.Connect()
	query := "SELECT idGallery, category, image, createdAt FROM gallery WHERE idGallery = ?"

	gallery := Gallery{}
	var createdAt time.Time

	err := con.QueryRow(query, idGallery).Scan(
		&gallery.IDGallery, &gallery.Category, &gallery.Image, &createdAt)

	gallery.CreatedAt = createdAt.Format("02 Jan 2006")

	defer con.Close()
	return gallery, err
}

// CreateGallery is func
func CreateGallery(data Gallery) error {
	con := db.Connect()
	query := "INSERT INTO gallery (image, category, createdAt) VALUES (?,?,?)"
	_, err := con.Exec(query, data.Image, data.Category, data.CreatedAt)

	defer con.Close()

	return err
}

// DeleteGallery is func
func DeleteGallery(idGallery string) error {
	con := db.Connect()
	query := "DELETE FROM gallery WHERE idGallery = ?"
	_, err := con.Exec(query, idGallery)

	defer con.Close()

	return err
}
