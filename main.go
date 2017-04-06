package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	sentence := speak("Duco")
	fmt.Printf("The parrot says %s\n", sentence)

	t, e := template.New("helloTemplate").Parse("Hello {{.}}\n")
	if e != nil {
		panic(e)
	}
	t.Execute(os.Stdout, "Quintor")
}

func speak(name string) string {
	return "Hello " + name
}
