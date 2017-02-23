package main

import (
"net/http"
"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {
	http.HandleFunc("/", name)
	http.ListenAndServe(":8080", nil)
}

func name(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "name.gohtml", "Aaron")
}