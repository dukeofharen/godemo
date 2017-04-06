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

type sumRequest struct {
	Start   int   `json:"start"`
	End     int   `json:"end"`
	Numbers []int `json:"numbers"`
}

type sumResponse struct {
	Answer       int `json:"answer"`
	Contributers int `json:"contributers"`
}

func main() {
	s := server{name: "Golang hands-on"}

	http.Handle("/", http.FileServer(http.Dir("./resources/")))
	http.HandleFunc("/api/greet/", s.titleHandler)
	http.HandleFunc("/api/sum/", s.sumHandler)

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

func (s server) sumHandler(w http.ResponseWriter, r *http.Request) {
	var sreq sumRequest
	d := json.NewDecoder(r.Body)
	err := d.Decode(&sreq)

	if err != nil {
		log.Println(err)
	}

	var answer int
	contributors := sreq.End - sreq.Start

	for _, element := range sreq.Numbers[sreq.Start:sreq.End] {
		answer += element
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")

	e := json.NewEncoder(w)
	err = e.Encode(sumResponse{Answer: answer, Contributers: contributors})
}

func speak(name string) string {
	return "Hello " + name
}
