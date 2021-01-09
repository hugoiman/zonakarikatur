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

// GetFrames is func
func GetFrames(w http.ResponseWriter, r *http.Request) {
	offset, _ := strconv.Atoi(r.FormValue("offset"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))

	data := models.GetFrames(offset, limit)
	message, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// CreateFrame is func
func CreateFrame(w http.ResponseWriter, r *http.Request) {
	var frame models.Frame
	if err := json.NewDecoder(r.Body).Decode(&frame); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err := validator.New().Struct(frame); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	frame.CreatedAt = time.Now().Format("2006-01-02")

	err := models.CreateFrame(frame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Gambar berhasil disimpan!"}`))
}

// DeleteFrame is func
func DeleteFrame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idFrame := vars["idFrame"]

	frame, err := models.GetFrame(idFrame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.DeleteFrame(idFrame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := "assets2/images/frame/" + frame.Image
	_ = os.Remove(path)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Frame telah dihapus!"}`))
}

// UploadFileFrame to server & cloudinary
func UploadFileFrame(w http.ResponseWriter, r *http.Request) {
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
			if outfile, err = os.Create("./assets2/images/frame/" + hdr.Filename); nil != err {
				status = http.StatusBadRequest
				return
			}
			// 32K buffer copy
			if _, err = io.Copy(outfile, infile); nil != err {
				status = http.StatusBadRequest
				return
			}

			publicID, err = cloudinary.GetService().UploadFile(strconv.Itoa(key), "zonakarikatur/frame", outfile.Name(), nil)
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
