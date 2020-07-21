package models

import "zonakarikatur/db"

// Admin is class
type Admin struct {
	IDAdmin  int    `json:"idAdmin"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required,min=3,max=10"`
	Email    string `json:"email" validate:"required,email"`
}

// GetAdmin is functio
func GetAdmin(id string) Admin {
	con := db.Connect()
	query := "SELECT idAdmin, name, username, email FROM admin WHERE idAdmin = ? OR email = ?"

	admin := Admin{}
	_ = con.QueryRow(query, id, id).Scan(
		&admin.IDAdmin, &admin.Name, &admin.Username, &admin.Email)

	defer con.Close()

	return admin
}

// UpdateAdmin is func
func UpdateAdmin(admin Admin) error {
	con := db.Connect()
	query := "UPDATE admin SET name = ?, username = ?, email = ? WHERE idAdmin = ?"
	_, err := con.Exec(query, admin.Name, admin.Username, admin.Email, admin.IDAdmin)

	defer con.Close()

	return err
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
