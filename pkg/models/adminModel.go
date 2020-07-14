package models

import "zonakarikatur/db"

// Admin is class
type Admin struct {
	IDAdmin  int    `json:"idAdmin"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// GetAdmin is function
func GetAdmin(idAdmin string) Admin {
	con := db.Connect()
	query := "SELECT idAdmin, name, username, email FROM admin where idAdmin = ?"

	admin := Admin{}
	_ = con.QueryRow(query, idAdmin).Scan(
		&admin.IDAdmin, &admin.Name, &admin.Username, &admin.Email)

	defer con.Close()

	return admin
}

// CheckOldPassword is Auth User
func CheckOldPassword(idAdmin int, password string) bool {
	var isAny bool
	con := db.Connect()
	query := "SELECT EXISTS (SELECT 1 FROM admin WHERE idAdmin = ? AND password = ?)"
	con.QueryRow(query, idAdmin, password).Scan(&isAny)

	defer con.Close()

	return isAny
}

// UpdatePassword is Edit Password
func UpdatePassword(idAdmin int, password string) {
	con := db.Connect()
	query := "UPDATE admin SET password = ? WHERE idAdmin = ?"
	con.Exec(query, password, idAdmin)

	defer con.Close()
}
