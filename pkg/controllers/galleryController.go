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
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	gallery, err := models.GetGallery(idGallery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.DeleteGallery(idGallery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := "assets2/images/gallery/" + gallery.Image
	_ = os.Remove(path)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Gambar telah dihapus!"}`))
}

// UploadFileGallery to server & cloudinary
func UploadFileGallery(w http.ResponseWriter, r *http.Request) {
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
	var linkFile []string
	var publicID string
	for _, fheaders := range r.MultipartForm.File {
		for key, hdr := range fheaders {
			var infile multipart.File
			if infile, err = hdr.Open(); nil != err {
				status = http.StatusBadRequest
				return
			}
			// open destination
			var outfile *os.File
			if outfile, err = os.Create("./assets2/images/gallery/" + hdr.Filename); nil != err {
				status = http.StatusBadRequest
				return
			}
			// 32K buffer copy
			if _, err = io.Copy(outfile, infile); nil != err {
				status = http.StatusBadRequest
				return
			}

			publicID, err = cloudinary.GetService().UploadFile(strconv.Itoa(key), "zonakarikatur/gallery", outfile.Name(), nil)
			if err != nil {
				log.Fatal(err)
			}
			linkFile = append(linkFile, publicID)
		}
	}

	message, _ := json.Marshal(linkFile)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// UploadFileGallery is func
// func UploadFileGallery(w http.ResponseWriter, r *http.Request) {
// 	basePath, _ := os.Getwd()
// 	reader, err := r.MultipartReader()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	var fileName []string
// 	for {
// 		part, err := reader.NextPart()
// 		if err == io.EOF {
// 			break
// 		}

// 		if _, err = os.Stat("assets2/images/gallery"); os.IsNotExist(err) {
// 			os.MkdirAll("assets2/images/gallery", os.ModePerm)
// 		}

// 		fileLocation := filepath.Join(basePath, "assets2/images/gallery", part.FileName())
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
// 		fileName = append(fileName, part.FileName())
// 	}
// 	message, err := json.Marshal(fileName)

// 	w.Header().Set("Content-type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(message)
// }
