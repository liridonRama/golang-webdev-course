package main

import "net/http"

func main() {

	// http.Handle("/pics", )
	// http.Handle(http.Dir("starting-files"))

	http.ListenAndServe(":8080", http.FileServer(http.Dir("public")))
}
