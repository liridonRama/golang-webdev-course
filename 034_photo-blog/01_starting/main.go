package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	uuid "github.com/satori/go.uuid"
	"github.com/teris-io/shortid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", handleUpload)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handleUpload(w http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(5 * 1024 * 1024)

	file, header, err := req.FormFile("file")
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	xF := strings.Split(header.Filename, ".")

	sID, err := shortid.Generate()
	if err != nil {
		panic(sID)
	}

	fN := sID + "." + xF[1]
	fmt.Println(fN)

	localFile, err := os.Create("./images/" + fN)
	if err != nil {

		log.Panicln(err)
	}
	defer localFile.Close()

	io.Copy(localFile, file)

	c, err := req.Cookie("session")
	if err != nil {
		log.Panicln(err)
	}
	xS := strings.Split(c.Value, "|")
	xS = append(xS, fN)
	c.Value = strings.Join(xS, "|")

	http.SetCookie(w, c)
	http.Redirect(w, req, "/", 302)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)

	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}
