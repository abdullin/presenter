package main

import (
	"log"
	"net/http"

	"github.com/abdullin/presenter/kbd"
)

func main() {
	static := http.FileServer(http.Dir("static"))
	http.HandleFunc("/press", h)
	http.Handle("/", static)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func h(w http.ResponseWriter, r *http.Request) {
	press := r.FormValue("key")
	log.Println(press)
	kbd.KeyDown(press)
	kbd.KeyUp(press)
}
