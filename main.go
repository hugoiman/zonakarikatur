package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	mw "zonakarikatur/middleware"
	"zonakarikatur/pkg/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	origins := handlers.AllowedOrigins([]string{"*"})

	auth := router.PathPrefix("").Subrouter()
	auth.Use(mw.AuthToken)

	auth.HandleFunc("/api/auth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Verified"))
	}).Methods("GET")
	router.HandleFunc("/api/loginAdmin", controllers.LoginAdmin).Methods("POST")
	// router.HandleFunc("/api/logout", controllers.Logout).Methods("GET")

	auth.HandleFunc("/api/admin", controllers.GetAdmin).Methods("GET")
	auth.HandleFunc("/api/admin", controllers.UpdateAdmin).Methods("PUT")

	router.HandleFunc("/api/gallery", controllers.GetGalleries).Methods("GET")
	auth.HandleFunc("/api/gallery", controllers.CreateGallery).Methods("POST")
	auth.HandleFunc("/api/gallery/{idGallery}", controllers.DeleteGallery).Methods("DELETE")
	auth.HandleFunc("/api/gallery-file", controllers.UploadFileGallery).Methods("POST")

	router.HandleFunc("/api/testimony", controllers.GetTestimonies).Methods("GET")
	auth.HandleFunc("/api/testimony", controllers.CreateTestimony).Methods("POST")
	auth.HandleFunc("/api/testimony/{idTestimony}", controllers.DeleteTestimony).Methods("DELETE")
	auth.HandleFunc("/api/testimony-file", controllers.UploadFileTestimony).Methods("POST")

	router.HandleFunc("/api/offer", controllers.GetOffers).Methods("GET")
	auth.HandleFunc("/api/offer", controllers.CreateOffer).Methods("POST")
	auth.HandleFunc("/api/offer/{idOffer}", controllers.DeleteOffer).Methods("DELETE")
	auth.HandleFunc("/api/offer-file", controllers.UploadFileOffer).Methods("POST")

	router.HandleFunc("/api/faq", controllers.GetFaqs).Methods("GET")
	auth.HandleFunc("/api/faq", controllers.CreateFaq).Methods("POST")
	auth.HandleFunc("/api/faq/{idFaq}", controllers.DeleteFaq).Methods("DELETE")

	router.HandleFunc("/api/link", controllers.GetLink).Methods("GET")
	auth.HandleFunc("/api/link/{idLink}", controllers.UpdateLink).Methods("PUT")

	router.HandleFunc("/api/about", controllers.GetAbout).Methods("GET")
	auth.HandleFunc("/api/about/{idAbout}", controllers.UpdateAbout).Methods("PUT")

	auth.HandleFunc("/api/password", controllers.ChangePassword).Methods("POST")
	router.HandleFunc("/api/password", controllers.ForgotPassword).Methods("PUT")

	router.HandleFunc("/api/frame", controllers.GetFrames).Methods("GET")
	auth.HandleFunc("/api/frame", controllers.CreateFrame).Methods("POST")
	auth.HandleFunc("/api/frame/{idFrame}", controllers.DeleteFrame).Methods("DELETE")
	auth.HandleFunc("/api/frame-file", controllers.UploadFileFrame).Methods("POST")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	router.PathPrefix("/assets2/").Handler(http.StripPrefix("/assets2/", http.FileServer(http.Dir("assets2"))))
	router.PathPrefix("/node_modules/").Handler(http.StripPrefix("/node_modules/", http.FileServer(http.Dir("node_modules2"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/index.html",
			"./pkg/views/header.html",
			"./pkg/views/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/gallery.html",
			"./pkg/views/header.html",
			"./pkg/views/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/bingkai", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/frame.html",
			"./pkg/views/header.html",
			"./pkg/views/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	//	Admin template
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/login.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/forgot-password", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/forgot-password.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/profile", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/profile.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/password", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/password.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/dashboard", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/dashboard.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/offer", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/offer.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/create-offer", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/offer-create.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/gallery", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/gallery.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/create-gallery", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/gallery-create.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/testimony", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/testimony.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/create-testimony", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/testimony-create.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/faq", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/faq.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/about", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/about.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/link", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/link-order.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/frame", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/frame.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	router.HandleFunc("/admin/create-frame", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.ParseFiles(
			"./pkg/views/admin/frame-create.html",
			"./pkg/views/admin/header.html",
			"./pkg/views/admin/footer.html",
		))

		var err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	// os.Setenv("PORT", "8080")
	// port := "8080"
	port := os.Getenv("PORT")

	fmt.Println("Server running at :", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(origins)(router)))
}
