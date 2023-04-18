package main

import (
	"errors"
	"fmt"
	"net/http"
	"html/template"
)

// TODO func registerTemplates(){}

func getHome(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./templates/base.html",
		"./templates/header.html",
		"./templates/index.html",
	}

	fmt.Printf("got / request\n")

	t, _ := template.ParseFiles(files...)
	t.ExecuteTemplate(w, "base", "data")
}

func getListings(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./templates/base.html",
		"./templates/header.html",
		"./templates/listings.html",
	}

	fmt.Printf("got /hello request\n")

	t, _ := template.ParseFiles(files...)
	t.ExecuteTemplate(w, "base", "data")
}

func main() {

	http.HandleFunc("/", getHome)
	http.HandleFunc("/listings", getListings)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed.\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}