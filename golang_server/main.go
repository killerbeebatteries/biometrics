package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
  "github.com/gorilla/schema"

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

	getAllBiometricData := func() ([]Biometric, error) {
		var biometrics []Biometric
		rows, err := DB.Query("SELECT id, date, time, sys, dia, bp, weight_total, weight_fat, weight_muscle, comment FROM bp_and_weight ORDER BY date DESC, time DESC")

		if err != nil {
      log.Print("error retrieving biometrics data from database: %v", err)
      return nil, err
		}

		defer rows.Close()

		for rows.Next() {

			var biometric Biometric

			err = rows.Scan(&biometric.Id, &biometric.Date, &biometric.Time, &biometric.Sys, &biometric.Dia, &biometric.Bp, &biometric.Weight_total, &biometric.Weight_fat, &biometric.Weight_muscle, &biometric.Comment)
			if err != nil {
        log.Print("error assigning biometrics data from database: %v", err)
        return nil, err
			}

			biometrics = append(biometrics, biometric)
		}

		return biometrics, nil
	}

  addBiometricData := func(biometric Biometric) error {
    _, err := DB.Exec("INSERT INTO bp_and_weight (date, time, sys, dia, bp, weight_total, weight_fat, weight_muscle, comment) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", biometric.Date, biometric.Time, biometric.Sys, biometric.Dia, biometric.Bp, biometric.Weight_total, biometric.Weight_fat, biometric.Weight_muscle, biometric.Comment)
    if err != nil {
      log.Fatal(err)
    }
    return nil
  }

  handleAddBiometric := func(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
      http.Error(w, "Failed to parse form values", http.StatusInternalServerError)
      return
    }

    var biometric Biometric
    decoder := schema.NewDecoder()

    err = decoder.Decode(&biometric, r.PostForm)

    if err != nil {
      http.Error(w, "Failed to decode form values", http.StatusInternalServerError)
      return
    }

    err = addBiometricData(biometric)
    if err != nil {
      http.Error(w, "Failed to add biometric data", http.StatusInternalServerError)
      return
    }

  }

	handleMainPage := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("test.html"))
		biometrics := getAllBiometricData()

		data := map[string][]Biometric{
			"Biometrics": biometrics,
		}

		tmpl.Execute(w, data)
	}

	// define handlers
	http.HandleFunc("/", handleMainPage)
	http.HandleFunc("/add-biometric", handleAddBiometric)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
