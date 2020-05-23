package main

import (
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", fs))

	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	templates, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		panic(err)
	}

	templates.ExecuteTemplate(res, "index.gohtml", nil)
}
