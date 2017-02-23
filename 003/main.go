package main

import (
  "html/template"
  "net/http"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandlerFunc("/", data)
	http.ListenAndServe(":8080", nil)
}

func data(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "data.gohtml", 256)
}
