package controllers

import (
	"encoding/json"
	"net/http"

	"zonakarikatur/pkg/models"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

// GetLink is func
func GetLink(w http.ResponseWriter, r *http.Request) {
	data := models.GetLink()
	message, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// UpdateLink is func
func UpdateLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idLink := vars["idLink"]

	var link models.Link
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err := validator.New().Struct(link); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := models.UpdateLink(idLink, link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Data berhasil disimpan!"}`))
}
