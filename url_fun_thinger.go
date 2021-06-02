package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

var quantity = []string{"two", "three", "four", "five", "six", "eight", "nine", "ten", "twelve"}
var adjectives = []string{"merry", "witty", "lovely", "sweet", "nice", "fine",
	"chill", "fresh"}
var nouns = []string{"cats", "dogs", "birds", "chairs", "trees", "books", "clouds",
	"fish"}

type URLSubmit struct {
	URLstr    string
	Shortened string
	Success   bool
	ErrorMsg  string
}

var default_URLSubmit = URLSubmit{
	URLstr:    "",
	Shortened: "",
	Success:   false,
	ErrorMsg:  "",
}

//go:embed templates/*
var templateData embed.FS

func persistShortened(shortened string, URLstr string) bool {
	// return 'false' if the 'shortened' is already in use

	jsonFileName := "persisted.json"

	jsonFile, err := os.Open(jsonFileName)
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var objmap map[string]interface{}
	if err := json.Unmarshal(byteValue, &objmap); err != nil {
		fmt.Println(err)
	}

	if _, ok := objmap[shortened]; !ok {
		objmap[shortened] = URLstr
		file, _ := json.MarshalIndent(objmap, "", "  ")
		_ = ioutil.WriteFile(jsonFileName, file, 0644)
		return true
	}

	// otherwise
	return false
}

func randomShortened() string {
	rand.Seed(time.Now().UnixNano())
	rQuantity := quantity[rand.Intn(len(quantity))]
	rAjective := adjectives[rand.Intn(len(adjectives))]
	rNoun := nouns[rand.Intn(len(nouns))]
	return rQuantity + rAjective + rNoun
}

func main() {

	tmpl, err := template.ParseFS(templateData, "templates/home.html")
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, default_URLSubmit)
			return
		}

		shortened := randomShortened()

		// Retry random string if already exists
		for written := persistShortened(shortened, r.FormValue("URLstr")); !written; written = persistShortened(shortened, r.FormValue("URLstr")) {
			fmt.Printf("colision detect: %v\n", shortened)
			shortened = randomShortened()
		}

		details := URLSubmit{
			URLstr:    r.FormValue("URLstr"),
			Shortened: shortened,
			Success:   true,
			ErrorMsg:  "",
		}

		tmpl.Execute(w, details)
	})

	http.HandleFunc("/lookup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			// (this should not be reachable)
			tmpl.Execute(w, default_URLSubmit)
			return
		} else {
			error_URLSubmit := URLSubmit{
				URLstr:    "",
				Shortened: "",
				Success:   false,
				ErrorMsg:  "We could not provide the code you supplied!!!",
			}

			tmpl.Execute(w, error_URLSubmit)
			return
		}
	})

	http.ListenAndServe(":8080", nil)

}
