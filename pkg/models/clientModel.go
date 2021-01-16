package models

import (
	"time"

	"zonakarikatur/db"
)

// Clients is func
type Clients struct {
	Clients []Client `json:"clients"`
}

// Client is func
type Client struct {
	IDClient  int    `json:"idClient"`
	Image     string `json:"image" validate:"required"`
	CreatedAt string `json:"createdAt"`
}

// GetClients is func
func GetClients() Clients {
	con := db.Connect()
	query := "SELECT idClient, image, createdAt FROM client"
	rows, _ := con.Query(query)

	var createdAt time.Time
	client := Client{}
	clients := Clients{}

	for rows.Next() {
		rows.Scan(
			&client.IDClient, &client.Image, &createdAt,
		)

		client.CreatedAt = createdAt.Format("02 Jan 2006")
		clients.Clients = append(clients.Clients, client)
	}

	defer con.Close()

	return clients
}

// GetClient is func
func GetClient(idClient string) (Client, error) {
	con := db.Connect()
	query := "SELECT idClient, image, createdAt FROM client WHERE idClient = ?"

	client := Client{}
	var createdAt time.Time

	err := con.QueryRow(query, idClient).Scan(
		&client.IDClient, &client.Image, &createdAt)

	client.CreatedAt = createdAt.Format("02 Jan 2006")

	defer con.Close()
	return client, err
}

// CreateClient is func
func CreateClient(data Client) error {
	con := db.Connect()
	query := "INSERT INTO client (image, createdAt) VALUES (?,?)"
	_, err := con.Exec(query, data.Image, data.CreatedAt)

	defer con.Close()

	return err
}

// DeleteClient is func
func DeleteClient(idClient string) error {
	con := db.Connect()
	query := "DELETE FROM client WHERE idClient = ?"
	_, err := con.Exec(query, idClient)

	defer con.Close()

	return err
}
