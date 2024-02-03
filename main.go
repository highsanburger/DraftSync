package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
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

	log.Println("App running on 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
