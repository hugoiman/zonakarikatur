package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"zonakarikatur/pkg/models"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

// GetOffers is func
func GetOffers(w http.ResponseWriter, r *http.Request) {
	data := models.GetOffers()
	message, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// CreateOffer is func
func CreateOffer(w http.ResponseWriter, r *http.Request) {
	var offer models.Offer
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err := validator.New().Struct(offer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offer.CreatedAt = time.Now().Format("2006-01-02")

	err := models.CreateOffer(offer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Data berhasil disimpan!"}`))
}

// DeleteOffer is func
func DeleteOffer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idOffer := vars["idOffer"]

	offer, err := models.GetOffer(idOffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.DeleteOffer(idOffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := "assets2/images/offer/" + offer.Image
	_ = os.Remove(path)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Gambar telah dihapus!"}`))
}

// UploadFileOffer is func
func UploadFileOffer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hai")

	basePath, _ := os.Getwd()
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := ""
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		fileLocation := filepath.Join(basePath, "assets2/images/offer", part.FileName())
		dst, err := os.Create(fileLocation)
		if dst != nil {
			defer dst.Close()
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fileName = part.FileName()
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "` + fileName + `"}`))
}
