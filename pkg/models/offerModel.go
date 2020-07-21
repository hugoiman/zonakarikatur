package models

import (
	"time"
	"zonakarikatur/db"
)

// Offer is class
type Offer struct {
	IDOffer   int    `json:"idOffer"`
	Image     string `json:"image" validate:"required"`
	Title     string `json:"title" validate:"required"`
	CreatedAt string `json:"createdAt"`
}

// Offers is offer list
type Offers struct {
	Offers []Offer `json:"offers"`
}

// GetOffers is func
func GetOffers() Offers {
	con := db.Connect()
	query := "SELECT idOffer, title, image, createdAt FROM offer"
	rows, _ := con.Query(query)

	var createdAt time.Time
	offer := Offer{}
	offers := Offers{}

	for rows.Next() {
		rows.Scan(
			&offer.IDOffer, &offer.Title, &offer.Image, &createdAt,
		)

		offer.CreatedAt = createdAt.Format("02 Jan 2006")
		offers.Offers = append(offers.Offers, offer)
	}

	defer con.Close()

	return offers
}

// GetOffer is func
func GetOffer(idOffer string) (Offer, error) {
	con := db.Connect()
	query := "SELECT idOffer, title, image, createdAt FROM offer WHERE idOffer = ?"

	offer := Offer{}
	var createdAt time.Time

	err := con.QueryRow(query, idOffer).Scan(
		&offer.IDOffer, &offer.Title, &offer.Image, &createdAt)

	offer.CreatedAt = createdAt.Format("02 Jan 2006")

	defer con.Close()
	return offer, err
}

// CreateOffer is func
func CreateOffer(data Offer) error {
	con := db.Connect()
	query := "INSERT INTO offer (image, title, createdAt) VALUES (?,?,?)"
	_, err := con.Exec(query, data.Image, data.Title, data.CreatedAt)

	defer con.Close()

	return err
}

// DeleteOffer is func
func DeleteOffer(idOffer string) error {
	con := db.Connect()
	query := "DELETE FROM offer WHERE idOffer = ?"
	_, err := con.Exec(query, idOffer)

	defer con.Close()

	return err
}
