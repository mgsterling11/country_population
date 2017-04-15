package main

import (
	"net/http"
	"net/url"

	"encoding/json"

	"strings"
	"time"
)

var replacer = strings.NewReplacer("_", " ")

func countryInformation(country string) (CountryPopulation, error) {
	country = strings.Title(replacer.Replace(country))
	encodedCountryName := &url.URL{Path: country}

	date := time.Now().Format("2006-01-02")

	resp, err := http.Get("http://api.population.io/1.0/population/" + encodedCountryName.String() + "/" + date)

  if err != nil {
    return CountryPopulation{}, err
  }

  defer resp.Body.Close()

  var population CountryPopulation

  if err := json.NewDecoder(resp.Body).Decode(&population); err != nil {
    return CountryPopulation{}, err
  }

  return population, nil
}

func serverConnect(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server is connected to port 8080"))
}

func main() {
	// tests local server is running on port 8080
	http.HandleFunc("/test-server", serverConnect)

	// retrieve country population
	http.HandleFunc("/country/", func(w http.ResponseWriter, r *http.Request) {
		country := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := countryInformation(country)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

type CountryPopulation struct {
	TotalPopulation struct {
		Date string `json:"date"`
		Population int `json:"population"`
	}`json:"total_population"`
}

// SAMPLE RESPONSE FROM http://api.population.io:80/1.0/population/Brazil/2015-12-24/
// {
//   "total_population": {
//     "date": "2015-12-24",
//     "population": 208679204
//   }
// }
