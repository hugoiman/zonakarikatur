package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"zonakarikatur/pkg/models"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

// GetFaqs is func
func GetFaqs(w http.ResponseWriter, r *http.Request) {
	data := models.GetFaqs()
	message, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// CreateFaq is func
func CreateFaq(w http.ResponseWriter, r *http.Request) {
	var faq models.Faq
	if err := json.NewDecoder(r.Body).Decode(&faq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err := validator.New().Struct(faq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idFaq, err := models.CreateFaq(faq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"idFaq":"` + strconv.Itoa(idFaq) + `"}`))
}

// DeleteFaq is func
func DeleteFaq(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idFaq := vars["idFaq"]

	_, err := models.GetFaq(idFaq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.DeleteFaq(idFaq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Data telah dihapus!"}`))
}
