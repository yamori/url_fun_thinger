package main

import (
	"embed"
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
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
}

//go:embed templates/*
var templateData embed.FS

func main() {

	tmpl, err := template.ParseFS(templateData, "templates/home.html")
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		rQuantity := quantity[rand.Intn(len(quantity))]
		rAjective := adjectives[rand.Intn(len(adjectives))]
		rNoun := nouns[rand.Intn(len(nouns))]

		details := URLSubmit{
			URLstr:    r.FormValue("URLstr"),
			Shortened: rQuantity + rAjective + rNoun,
			Success:   true,
		}

		// do something with details
		_ = details
		fmt.Println(details)

		tmpl.Execute(w, details)
	})

	http.ListenAndServe(":8080", nil)

}
