package server

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func Start() {
	fmt.Println("server started")
	fmt.Println("http://localhost:8080")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.tmpl"))

	value, err := loadMd()
	if err != nil {
		panic(err)
	}

	m := map[string]string{
		"Erd": value,
	}

	tmpl.Execute(w, m)
}

func loadMd() (string, error) {
	f, err := os.Open("mermaid.md")
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
