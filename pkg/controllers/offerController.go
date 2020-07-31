package controllers

import (
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"zonakarikatur/internal/cloudinary"
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

	// offer, err := models.GetOffer(idOffer)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	err := models.DeleteOffer(idOffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// path := "assets2/images/offer/" + offer.Image
	// _ = os.Remove(path)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Gambar telah dihapus!"}`))
}

// Upload to server & cloudinary
func UploadFileOffer(w http.ResponseWriter, r *http.Request) {
	var (
		status int
		err    error
	)
	defer func() {
		if nil != err {
			http.Error(w, err.Error(), status)
		}
	}()

	const _24K = (1 << 20) * 24
	if err = r.ParseMultipartForm(_24K); nil != err {
		status = http.StatusBadRequest
		return
	}
	var publicId string
	for _, fheaders := range r.MultipartForm.File {
		for key, hdr := range fheaders {
			var infile multipart.File
			if infile, err = hdr.Open(); nil != err {
				status = http.StatusBadRequest
				return
			}
			// open destination
			var outfile *os.File
			if outfile, err = os.Create("./assets2/images/offer/" + hdr.Filename); nil != err {
				status = http.StatusBadRequest
				return
			}
			// 32K buffer copy
			if _, err = io.Copy(outfile, infile); nil != err {
				status = http.StatusBadRequest
				return
			}

			publicId, err = cloudinary.GetService().UploadFile(strconv.Itoa(key), "zonakarikatur/offer", outfile.Name(), nil)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	message, _ := json.Marshal(publicId)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// Upload to Server
// func UploadFileOffer2(w http.ResponseWriter, r *http.Request) {
// 	basePath, _ := os.Getwd()
// 	reader, err := r.MultipartReader()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	fileName := ""
// 	for {
// 		part, err := reader.NextPart()
// 		if err == io.EOF {
// 			break
// 		}

// 		if _, err = os.Stat("assets2/images/offer"); os.IsNotExist(err) {
// 			os.MkdirAll("assets2/images/offer", os.ModePerm)
// 		}

// 		fileLocation := filepath.Join(basePath, "assets2/images/offer", part.FileName())
// 		dst, err := os.Create(fileLocation)
// 		if dst != nil {
// 			defer dst.Close()
// 		}
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		if _, err := io.Copy(dst, part); err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		fileName = part.FileName()
// 	}

// 	w.Header().Set("Content-type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "` + fileName + `"}`))
// }
