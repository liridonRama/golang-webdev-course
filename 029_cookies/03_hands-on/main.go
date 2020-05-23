package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})

	v, err := req.Cookie("visits")

	if err != nil {
		fmt.Println("damn")
		http.SetCookie(w, &http.Cookie{
			Name:  "visits",
			Value: "1",
		})
	} else {
		y, err := strconv.Atoi(v.Value)
		if err != nil {
			log.Fatalln("err while parsing visits cookie")
		}
		y++

		v.Value = strconv.Itoa(y)
		http.SetCookie(w, v)
	}

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	} else {
		c.MaxAge = -1
		http.SetCookie(w, c)
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
