package player

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := &Page{Title: "live stream player"}

	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func New() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8099", nil)
}
