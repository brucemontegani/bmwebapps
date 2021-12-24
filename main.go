package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTmpl, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		log.Println("error parsing template ", tmpl)
		log.Println(err)
	}

	if err := parsedTmpl.Execute(w, nil); err != nil {
		log.Println("error executing template ", tmpl)
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting application on port", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Println(err)
	}
}
