package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type server struct {
	name string
}

type greet struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

func main() {
	s := server{name: "Golang hands-on"}

	http.Handle("/", http.FileServer(http.Dir("./resources/")))
	http.HandleFunc("/api/greet/", s.titleHandler)

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func (s server) titleHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	accept := r.Header.Get("Accept")

	if accept == "application/json" {
		w.Header().Set("Content-Type", "application/json; charset=utf8")

		e := json.NewEncoder(w)
		err = e.Encode(greet{Title: s.name, Subtitle: "Hello Quintor!"})
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=utf8")

		_, err = fmt.Fprint(w, s.name)
	}

	if err != nil {
		log.Println(err)
	}
}

func speak(name string) string {
	return "Hello " + name
}
