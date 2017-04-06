package main

import (
	"log"
	"net/http"
	"text/template"
)

type server struct {
	template *template.Template
}

func main() {
	t := template.Must(template.ParseFiles("templates/index.html"))
	s := server{template: t}

	http.HandleFunc("/", s.handleRoot)
	http.Handle("/resources/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)

	// sentence := speak("Duco")
	// fmt.Printf("The parrot says %s\n", sentence)

	// t, e := template.ParseFiles("template.tpl")
	// if e != nil {
	// 	panic(e)
	// }

	// t.Execute(os.Stdout, "Quintor")
}

func (s *server) handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")

	err := s.template.Execute(w, "Vrolijk pasen!")

	if err != nil {
		log.Println(err)
	}
}

func speak(name string) string {
	return "Hello " + name
}
