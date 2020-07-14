package controllers

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"

	models "zonakarikatur/pkg/models"

	"net/http"

	"github.com/gorilla/context"
	"gopkg.in/go-playground/validator.v9"
)

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
