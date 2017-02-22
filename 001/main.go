package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Index</title>
</head>
<body>
Hello from index
<a href="/index">index</a>
<a href="/about">about</a>
<a href="/contact">contact</a>
</body>
</html>`)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>about</title>
</head>
<body>
Hello from about
<a href="/index">index</a>
<a href="/about">about</a>
<a href="/contact">contact</a>
</body>
</html>`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>contact</title>
</head>
<body>
Hello from contact
<a href="/index">index</a>
<a href="/about">about</a>
<a href="/contact">contact</a>
</body>
</html>`)
}