package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Biometric struct {
	date          time.Time
	time          time.Time
	sys           int
	dia           int
	bp            int
	weight_total  float64
	weight_fat    float64
	weight_muscle float64
	comment       string
}

func main() {
	fmt.Println("Go app...")

	// handler function #1 - returns the index.html template, with film data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		biometrics := map[string][]Biometric{
			"Biometrics": {
				{date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), sys: 120, dia: 80, bp: 80, weight_total: 80.0, weight_fat: 20.0, weight_muscle: 60.0, comment: "comment 1"},
				{date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), sys: 120, dia: 80, bp: 80, weight_total: 80.0, weight_fat: 20.0, weight_muscle: 60.0, comment: "comment 1"},
				{date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), sys: 120, dia: 80, bp: 80, weight_total: 80.0, weight_fat: 20.0, weight_muscle: 60.0, comment: "comment 1"},
			},
		}
		tmpl.Execute(w, biometrics)
	}

	// handler function #2 - returns the template block with the newly added biometrics, as an HTMX response
	h2 := func(w http.ResponseWriter, r *http.Request) {
		date, err := time.Parse("2006-01-02", r.PostFormValue("date"))
		if err != nil {
			log.Fatal(err)
		}
		time, err := time.Parse("15:04", r.PostFormValue("time"))
		if err != nil {
			log.Fatal(err)
		}
		sys, err := strconv.Atoi(r.PostFormValue("sys"))
		if err != nil {
			log.Fatal(err)
		}
		dia, err := strconv.Atoi(r.PostFormValue("dia"))
		if err != nil {
			log.Fatal(err)
		}
		bp, err := strconv.Atoi(r.PostFormValue("bp"))
		if err != nil {
			log.Fatal(err)
		}
		weight_total, err := strconv.ParseFloat(r.PostFormValue("weight_total"), 64)
		if err != nil {
			log.Fatal(err)
		}
		weight_fat, err := strconv.ParseFloat(r.PostFormValue("weight_fat"), 64)
		if err != nil {
			log.Fatal(err)
		}
		weight_muscle, err := strconv.ParseFloat(r.PostFormValue("weight_muscle"), 64)
		if err != nil {
			log.Fatal(err)
		}
		comment := r.PostFormValue("comment")

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "biometrics-list-element", Biometric{date: date, time: time, sys: sys, dia: dia, bp: bp, weight_total: weight_total, weight_fat: weight_fat, weight_muscle: weight_muscle, comment: comment})
	}

	// define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
