package main

import (
	"log"
	"net/http"

	"github.com/abdullin/presenter/kbd"
)

func main() {
	http.HandleFunc("/press", h)
	static :=  http.FileServer(http.Dir("static"))
	http.Handle("/", static)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func h(w http.ResponseWriter, r *http.Request) {
	press := r.FormValue("key")
	kbd.KeyPress(press)
}
