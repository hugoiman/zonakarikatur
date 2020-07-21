package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"zonakarikatur/pkg/models"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

// GetTestimonies is func
func GetTestimonies(w http.ResponseWriter, r *http.Request) {
	offset, _ := strconv.Atoi(r.FormValue("offset"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))

	data := models.GetTestimonies(offset, limit)
	message, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// CreateTestimony is func
func CreateTestimony(w http.ResponseWriter, r *http.Request) {
	var testimony models.Testimony
	if err := json.NewDecoder(r.Body).Decode(&testimony); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err := validator.New().Struct(testimony); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	testimony.CreatedAt = time.Now().Format("2006-01-02")

	err := models.CreateTestimony(testimony)
	if err != nil {
		http.Error(w, "Gagal! Nama image sudah ada.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Testimony berhasil disimpan!"}`))
}

// DeleteTestimony is func
func DeleteTestimony(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idTestimony := vars["idTestimony"]

	testimony, err := models.GetTestimony(idTestimony)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.DeleteTestimony(idTestimony)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := "assets2/images/gallery/" + testimony.Image
	_ = os.Remove(path)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Gambar telah dihapus!"}`))
}
