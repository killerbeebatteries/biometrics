package main

import (
	"fmt"
	"github.com/gorilla/schema"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"regexp"
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

	// addBiometricData := func(biometric Biometric) error {
	//   _, err := DB.Exec("INSERT INTO bp_and_weight (date, time, sys, dia, bp, weight_total, weight_fat, weight_muscle, comment) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", biometric.Date, biometric.Time, biometric.Sys, biometric.Dia, biometric.Bp, biometric.Weight_total, biometric.Weight_fat, biometric.Weight_muscle, biometric.Comment)
	//   if err != nil {
	//     log.Print("error inserting database record: %v", err)
	//     return err
	//   }
	//   return nil
	// }

	handleAddBiometric := func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Print("Failed to parse form from http request: %v", err)
			http.Error(w, "Failed to parse form values", http.StatusInternalServerError)
			return
		}

		// cool work-around for dealing with date types from the form
		var timeConverter = func(value string) reflect.Value {

			datePattern := `^\d{4}-\d{2}-\d{2}$`
			r := regexp.MustCompile(datePattern)

			if r.MatchString(value) {
				if v, err := time.Parse("2006-01-02", value); err == nil {
					return reflect.ValueOf(v)
				}
			}

      timePattern := `^([01]\d|2[0-3]):[0-5]\d$`
			r = regexp.MustCompile(timePattern)
			if r.MatchString(value) {
				if v, err := time.Parse("15:04", value); err == nil {
					return reflect.ValueOf(v)
				}
			}

			return reflect.Value{}
		}

		var biometric Biometric
		decoder := schema.NewDecoder()

		decoder.RegisterConverter(time.Time{}, timeConverter)

		err = decoder.Decode(&biometric, r.PostForm)

		fmt.Printf("Biometric: %v", biometric)

		if err != nil {
			log.Print("Failed to decode form values: %v", err)
			http.Error(w, "Failed to decode form values", http.StatusInternalServerError)
			return
		}

		// err = addBiometricData(biometric)
		// if err != nil {
		//   log.Print("Failed to add biometric data: %v", err)
		//   http.Error(w, "Failed to add biometric data", http.StatusInternalServerError)
		//   return
		// }
		//
		// // redirect to main page
		// http.Redirect(w, r, "/", http.StatusFound)

	}

	handleMainPage := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		biometrics, err := getAllBiometricData()

		if err != nil {
			log.Print("Failed to get biometric data: %v", err)
			http.Error(w, "Failed to get biometric data", http.StatusInternalServerError)
			return
		}

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
