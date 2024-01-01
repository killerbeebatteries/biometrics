package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)
// using pointers in the data structure to support null values
// assigned from the database query results.
type Biometric struct {
  Id            int
	Date          time.Time
	Time          *time.Time
	Sys           *int
	Dia           *int
	Bp            *int
	Weight_total  *float64
	Weight_fat    *float64
	Weight_muscle *float64
	Comment       string
}

func main() {
	fmt.Println("Go Biometrics app...")

  err := OpenDatabase()
	if err != nil {
		log.Fatal(err)
	}

  getAllWeightAndBPData := func() []Biometric {
    var biometrics []Biometric
    rows, err := DB.Query("SELECT id, date, time, sys, dia, bp, weight_total, weight_fat, weight_muscle, comment FROM bp_and_weight ORDER BY date DESC, time DESC")


    if err != nil {
      log.Fatal(err)
    }

    defer rows.Close()

    for rows.Next() {

      var biometric Biometric

      err = rows.Scan(&biometric.Id, &biometric.Date, &biometric.Time, &biometric.Sys, &biometric.Dia, &biometric.Bp, &biometric.Weight_total, &biometric.Weight_fat, &biometric.Weight_muscle, &biometric.Comment)
      if err != nil {
        log.Fatal(err)
      }

      biometrics = append(biometrics, biometric)
    }

    return biometrics
  }


  handleMainPage := func(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("test.html"))
    biometrics := getAllWeightAndBPData()

    data := map[string][]Biometric{
      "Biometrics": biometrics,
    }

    tmpl.Execute(w, data)
  }


	// define handlers
	http.HandleFunc("/", handleMainPage)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
