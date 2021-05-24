package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type URLSubmit struct {
	URLstr string
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := URLSubmit{
			URLstr: r.FormValue("URLstr"),
		}

		// do something with details
		_ = details
		fmt.Println(details)

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}
