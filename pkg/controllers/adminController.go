package controllers

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strconv"

	models "zonakarikatur/pkg/models"

	"net/http"

	"github.com/gorilla/context"
	"gopkg.in/go-playground/validator.v9"
)

// GetAdmin is func
func GetAdmin(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "admin").(*MyClaims)
	admin := models.GetAdmin(strconv.Itoa(user.IDAdmin))
	if admin.Name == "" {
		http.Error(w, "User tidak ditemukan.", http.StatusBadRequest)
		return
	}
	message, err := json.Marshal(admin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// UpdateAdmin is func
func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "admin").(*MyClaims)
	var admin models.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err := validator.New().Struct(admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	admin.IDAdmin = user.IDAdmin
	err := models.UpdateAdmin(admin)
	if err != nil {
		http.Error(w, "Gagal menyimpan perubahan. Coba lagi.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Password berhasil diperbarui!"}`))
}

// ChangePassword is Edit Password
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	admin := context.Get(r, "admin").(*MyClaims)

	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)

	if err := validator.New().Var(fmt.Sprintf("%v", data["password_baru"]), "required,min=6,max=18"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err := validator.New().Var(fmt.Sprintf("%v", data["password_lama"]), "required,min=6,max=18"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var oldPass = sha1.New()
	oldPass.Write([]byte(fmt.Sprintf("%v", data["password_lama"])))
	var encryptedOldPass = fmt.Sprintf("%x", oldPass.Sum(nil))

	isValid := models.CheckOldPassword(admin.IDAdmin, encryptedOldPass)
	if !isValid {
		http.Error(w, "Password lama tidak sesuai", http.StatusBadRequest)
		return
	}

	var newPass = sha1.New()
	newPass.Write([]byte(fmt.Sprintf("%v", data["password_baru"])))
	var encryptedPass = fmt.Sprintf("%x", newPass.Sum(nil))

	models.UpdatePassword(admin.IDAdmin, encryptedPass)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Password berhasil diperbarui!"}`))
}
