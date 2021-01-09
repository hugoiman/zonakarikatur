package models

import (
	"fmt"
	"time"

	"zonakarikatur/db"
)

// Frames is func
type Frames struct {
	Frames []Frame `json:"frames"`
}

// Frame is func
type Frame struct {
	IDFrame   int    `json:"idFrame"`
	Image     string `json:"image" validate:"required"`
	Model     string `json:"model" validate:"required"`
	CreatedAt string `json:"createdAt"`
}

// GetFrames is func
func GetFrames(limit, offset int) Frames {
	fmt.Println(limit, offset)
	con := db.Connect()
	query := "SELECT idFrame, image, model, createdAt FROM frame ORDER BY idFrame DESC LIMIT ?,?"
	rows, _ := con.Query(query, limit, offset)

	var createdAt time.Time
	frame := Frame{}
	frames := Frames{}

	for rows.Next() {
		rows.Scan(
			&frame.IDFrame, &frame.Image, &frame.Model, &createdAt,
		)

		frame.CreatedAt = createdAt.Format("02 Jan 2006")
		frames.Frames = append(frames.Frames, frame)
	}

	defer con.Close()

	return frames
}

// GetFrame is func
func GetFrame(idFrame string) (Frame, error) {
	con := db.Connect()
	query := "SELECT idFrame, model, image, createdAt FROM frame WHERE idFrame = ?"

	frame := Frame{}
	var createdAt time.Time

	err := con.QueryRow(query, idFrame).Scan(
		&frame.IDFrame, &frame.Model, &frame.Image, &createdAt)

	frame.CreatedAt = createdAt.Format("02 Jan 2006")

	defer con.Close()
	return frame, err
}

// CreateFrame is func
func CreateFrame(data Frame) error {
	con := db.Connect()
	query := "INSERT INTO frame (image, model, createdAt) VALUES (?,?,?)"
	_, err := con.Exec(query, data.Image, data.Model, data.CreatedAt)

	defer con.Close()

	return err
}

// DeleteFrame is func
func DeleteFrame(idFrame string) error {
	con := db.Connect()
	query := "DELETE FROM frame WHERE idFrame = ?"
	_, err := con.Exec(query, idFrame)

	defer con.Close()

	return err
}
