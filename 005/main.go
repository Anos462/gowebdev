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
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	m := map[string]bool{}
	m["Aaron"] = true
	m["Lilly"] = true
	tpl.ExecuteTemplate(w, "name.gohtml", m)
}