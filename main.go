package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type server struct {
    name string
    db   *gorm.DB
}
    
type Message struct {
    ID      uint
    Name    string
    Content string
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

type dbRequest struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type dbResponse struct {
	Id int `json:"id"`
}

func main() {
	db, err := gorm.Open("mysql", "root:geheim@/godemo")
	if err != nil {
		log.Fatal("Could not open database", err)
	}
	defer db.Close()
		
	db.DropTableIfExists(&Message{})
	db.AutoMigrate(&Message{})
		
	s := server{name: "Golang Hands-on", db: db}

	http.Handle("/", http.FileServer(http.Dir("./resources/")))
	http.HandleFunc("/api/greet/", s.titleHandler)
	http.HandleFunc("/api/sum/", s.sumHandler)
	http.HandleFunc("/api/store/", s.dbHandler)

	err = http.ListenAndServe(":8080", nil)
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

	if err != nil {
		log.Println(err)
	}
}

func (s server) dbHandler(w http.ResponseWriter, r *http.Request) {
	var dreq dbRequest
	d := json.NewDecoder(r.Body)
	d.Decode(&dreq)

	m := Message{Name: dreq.Name, Content: dreq.Message}
	s.db.Create(&m)
	log.Println("message id", m.ID)

	w.Header().Set("Content-Type", "application/json; charset=utf8")

	e := json.NewEncoder(w)
	err := e.Encode(dbResponse{Id: int(m.ID)})

	if err != nil {
		log.Println(err)
	}
}

func speak(name string) string {
	return "Hello " + name
}
