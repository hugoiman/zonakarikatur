package models

import "zonakarikatur/db"

// Auth is class
type Auth struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginAdmin is func
func LoginAdmin(login Auth) (string, error) {
	var idAdmin string

	con := db.Connect()
	query := "SELECT idAdmin FROM admin WHERE username = ? AND password = ?"
	err := con.QueryRow(query, login.Username, login.Password).Scan(&idAdmin)

	defer con.Close()
	return idAdmin, err

}
