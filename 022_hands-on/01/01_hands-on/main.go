package main

import (
	"io"
	"net/http"
)

func handleStandardRoute(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Im the standard Path")
}
func handleDogRoute(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "You Reached the dog route, Wouf Wouf!")
}
func handleMeRoute(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi, I'm Liridon Rama")
}

func main() {
	http.HandleFunc("/", handleStandardRoute)
	http.HandleFunc("/dog/", handleDogRoute)
	http.HandleFunc("/me/", handleMeRoute)

	http.ListenAndServe(":8080", nil)
}
