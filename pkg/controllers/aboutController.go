package controllers

import (
	"encoding/json"
	"net/http"

	"zonakarikatur/pkg/models"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

// GetAbout is func
func GetAbout(w http.ResponseWriter, r *http.Request) {
	data := models.GetAbout()
	message, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// UpdateAbout is func
func UpdateAbout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idAbout := vars["idLink"]

	var about models.About
	if err := json.NewDecoder(r.Body).Decode(&about); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err := validator.New().Struct(about); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := models.UpdateAbout(idAbout, about)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Data berhasil disimpan!"}`))
}
