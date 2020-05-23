package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpeg", pic)

	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	templates, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Panic(err)
	}

	templates.ExecuteTemplate(res, "dog.gohtml", "")
}

func pic(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./dog.jpeg")
}
