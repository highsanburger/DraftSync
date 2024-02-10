package main

import (
	"fmt"
	"html/template"
	"log"
	// "math/rand"
	"github.com/schollz/peerdiscovery"
	"net/http"
)

// type person struct {
//     name string
//     age  int
// }

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

	discoveries, _ := peerdiscovery.Discover(peerdiscovery.Settings{Limit: 1})
	for _, d := range discoveries {
		fmt.Printf("discovered '%s'\n", d.Address)
	}
}
