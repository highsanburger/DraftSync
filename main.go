package main

import (
	// "fmt"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		deviceName := r.FormValue("deviceName")
		tmpl := template.Must(template.ParseFiles("./templates/device_name.html"))
		tmpl.Execute(w, deviceName)
	})
	//
	// http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
	// 	file := r.FormValue("file")
	// 	tmpl := template.Must(template.ParseFiles("./templates/file_upload.html"))
	// 	tmpl.Execute(w, file)
	// 	fmt.Println(file)
	// })

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		// Parse the multipart form data
		err := r.ParseMultipartForm(10 << 20) // 10 MB limit for the file
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		fmt.Println(err)

		// Get the file from the form data
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Unable to get file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Save the file to the server
		fileName := handler.Filename
		filePath := filepath.Join("uploads", fileName) // Assuming you have a folder named "uploads"
		out, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Unable to create file on server", http.StatusInternalServerError)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "Unable to save file", http.StatusInternalServerError)
			return
		}

		// Respond with a success message
		w.Write([]byte("File uploaded successfully!"))
	})

	log.Println("App running on 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
