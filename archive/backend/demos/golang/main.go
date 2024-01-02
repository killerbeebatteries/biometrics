package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Biometric struct {
	sys   int
	dia   int
	pulse int
}

func main() {
	fmt.Println("Hello World")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

    biometrics := map[string][]Biometric{
      "Biometrics": {
        {sys: 120, dia: 80, pulse: 60},
        {sys: 122, dia: 83, pulse: 67},
        {sys: 123, dia: 85, pulse: 68},
      },
    }
		tmpl.Execute(w, biometrics)
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
