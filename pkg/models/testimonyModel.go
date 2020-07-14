package models

import (
	"time"

	"zonakarikatur/db"
)

// Testimonies is func
type Testimonies struct {
	Testimonies []Testimony `json:"testimonies"`
}

// Testimony is func
type Testimony struct {
	IDTestimony int    `json:"idTestimony"`
	Image       string `json:"image" validate:"required"`
	CreatedAt   string `json:"createdAt"`
}

// GetTestimonies is func
func GetTestimonies(limit, offset int) Testimonies {
	con := db.Connect()
	query := "SELECT idTestimony, image, createdAt FROM testimony ORDER BY idTestimony DESC LIMIT ?,?"
	rows, _ := con.Query(query, limit, offset)

	var createdAt time.Time
	testimony := Testimony{}
	testimonies := Testimonies{}

	for rows.Next() {
		rows.Scan(
			&testimony.IDTestimony, &testimony.Image, &createdAt,
		)

		testimony.CreatedAt = createdAt.Format("02 Jan 2006")
		testimonies.Testimonies = append(testimonies.Testimonies, testimony)
	}

	defer con.Close()

	return testimonies
}

// CreateTestimony is func
func CreateTestimony(data Testimony) error {
	con := db.Connect()
	query := "INSERT INTO testimony (image, createdAt) VALUES (?,?)"
	_, err := con.Exec(query, data.Image, data.CreatedAt)

	defer con.Close()

	return err
}

// DeleteTestimony is func
func DeleteTestimony(idTestimony string) error {
	con := db.Connect()
	query := "DELETE FROM testimony WHERE idTestimony = ?"
	_, err := con.Exec(query, idTestimony)

	defer con.Close()

	return err
}
