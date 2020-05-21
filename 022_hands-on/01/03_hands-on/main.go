package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func handleStandardRoute(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Im the standard Path")
}
func handleDogRoute(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "You Reached the dog route, Wouf Wouf!")
}
func handleMeRoute(res http.ResponseWriter, req *http.Request) {
	name := "Liridon Rama"

	t, err := template.ParseFiles("stuff.html")

	if err != nil {
		log.Panic(err)
	}

	err = t.ExecuteTemplate(res, "stuff.html", name)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/", handleStandardRoute)
	http.HandleFunc("/dog/", handleDogRoute)
	http.HandleFunc("/me/", handleMeRoute)

	http.ListenAndServe(":8080", nil)
}
