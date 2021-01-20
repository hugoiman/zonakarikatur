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

// GetTestimonies is func
func GetTestimonies(w http.ResponseWriter, r *http.Request) {
	// offset, _ := strconv.Atoi(r.FormValue("offset"))
	// limit, _ := strconv.Atoi(r.FormValue("limit"))

	data := models.GetTestimonies()
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

	// testimony, err := models.GetTestimony(idTestimony)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	err := models.DeleteTestimony(idTestimony)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// path := "assets2/images/testimony/" + testimony.Image
	// _ = os.Remove(path)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Gambar telah dihapus!"}`))
}

// UploadFileTestimony to server & cloudinary
func UploadFileTestimony(w http.ResponseWriter, r *http.Request) {
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
			if outfile, err = os.Create("./assets2/images/testimony/" + hdr.Filename); nil != err {
				status = http.StatusBadRequest
				return
			}
			// 32K buffer copy
			if _, err = io.Copy(outfile, infile); nil != err {
				status = http.StatusBadRequest
				return
			}

			publicID, err = cloudinary.GetService().UploadFile(strconv.Itoa(key), "zonakarikatur/testimony", outfile.Name(), nil)
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

// UploadFileTestimony is func
// func UploadFileTestimony(w http.ResponseWriter, r *http.Request) {
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

// 		if _, err = os.Stat("assets2/images/testimony"); os.IsNotExist(err) {
// 			os.MkdirAll("assets2/images/testimony", os.ModePerm)
// 		}

// 		fileLocation := filepath.Join(basePath, "assets2/images/testimony", part.FileName())
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
