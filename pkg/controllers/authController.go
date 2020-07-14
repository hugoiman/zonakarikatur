package controllers

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	mw "zonakarikatur/middleware"
	models "zonakarikatur/pkg/models"

	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/go-playground/validator.v9"
)

// MySigningKey is signature
var MySigningKey = mw.MySigningKey

// LoginAdmin is func
func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var login models.Auth
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err := validator.New().Struct(login); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var sha = sha1.New()
	sha.Write([]byte(login.Password))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	login.Password = encryptedString

	idAdmin, err := models.LoginAdmin(login)
	if err != nil {
		http.Error(w, "Gagal! Username atau password salah.", http.StatusBadRequest)
		return
	}

	admin := models.GetAdmin(idAdmin)
	token := CreateToken(admin)

	type M map[string]interface{}
	message, _ := json.Marshal(M{"token": token})

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// MyClaims is credential
type MyClaims = mw.MyClaims

// CreateToken is Generate token
func CreateToken(admin models.Admin) string {
	claims := MyClaims{
		IDAdmin:  admin.IDAdmin,
		Username: admin.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(MySigningKey)

	return tokenString
}

// Logout is func
// func Logout(w http.ResponseWriter, r *http.Request) {
// 	cookie := &http.Cookie{
// 		Name:    "tokenCookie",
// 		Value:   "",
// 		Path:    "/",
// 		Expires: time.Unix(0, 0),
// 	}
// 	http.SetCookie(w, cookie)

// 	w.Header().Set("Content-type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message":"Anda telah logout"}`))
// 	return
// }
