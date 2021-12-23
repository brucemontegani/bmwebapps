package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>This is the home page.</h1>")
}

func main() {
	http.HandleFunc("/", Home)

	fmt.Println("Starting application on port", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Println(err)
	}
}
