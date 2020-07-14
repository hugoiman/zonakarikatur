package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"zonakarikatur/pkg/models"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

// GetGalleries is func
func GetGalleries(w http.ResponseWriter, r *http.Request) {
	offset, _ := strconv.Atoi(r.FormValue("offset"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))

	data := models.GetGalleries(offset, limit)
	message, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// CreateGallery is func
func CreateGallery(w http.ResponseWriter, r *http.Request) {
	var gallery models.Gallery
	if err := json.NewDecoder(r.Body).Decode(&gallery); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err := validator.New().Struct(gallery); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gallery.CreatedAt = time.Now().Format("2006-01-02")

	err := models.CreateGallery(gallery)
	if err != nil {
		http.Error(w, "Gagal! Nama image sudah ada.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Gambar berhasil disimpan!"}`))
}

// DeleteGallery is func
func DeleteGallery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idGallery := vars["idGallery"]

	err := models.DeleteGallery(idGallery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Gambar telah dihapus!"}`))
}
