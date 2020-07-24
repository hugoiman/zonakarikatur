package controllers

import (
	"crypto/sha1"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	mw "zonakarikatur/middleware"
	models "zonakarikatur/pkg/models"

	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/gomail.v2"
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
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
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

// ForgotPassword is func
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)

	email := fmt.Sprintf("%v", data["email"])

	if err := validator.New().Var(email, "required,email"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	admin := models.GetAdmin(email)
	if admin.Name == "" {
		http.Error(w, "User tidak ditemukan", http.StatusBadRequest)
		return
	}

	// Generate Random String
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randomStr := make([]rune, 10)
	for i := range randomStr {
		randomStr[i] = letters[rand.Intn(len(letters))]
	}
	newPass := string(randomStr)

	var pass = sha1.New()
	pass.Write([]byte(newPass))
	var encryptedPass = fmt.Sprintf("%x", pass.Sum(nil))

	models.UpdatePassword(admin.IDAdmin, encryptedPass)
	message := "Hallo " + admin.Name + ", your new password is <b>" + newPass + "</b>"
	err := SendEmail("New Password", email, message)
	if err != nil {
		http.Error(w, "Gagal! Coba beberapa saat lagi.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"New password has been sent to your email."}`))
}

// SendEmail is func
func SendEmail(subject string, to string, message string) error {
	var configSMTPHost = "smtp.gmail.com"
	var configSMTPPort = 587
	var configEmail = "nanonymoux@gmail.com"
	var configPassword = os.Getenv("PASS_EMAIL")
	// var configPassword = "bkl"

	mailer := gomail.NewMessage()

	mailer.SetHeader("From", configEmail)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(configSMTPHost, configSMTPPort, configEmail, configPassword)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := dialer.DialAndSend(mailer)

	return err

}
