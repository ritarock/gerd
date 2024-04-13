package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

const TEMPATE = "internal/server/index.tmpl"
const MD = "mermaid.md"

func Start() {
	fmt.Println("server started")
	fmt.Println("http://localhost:8080")
	http.HandleFunc("/", indexHander)
	http.ListenAndServe(":8080", nil)
}

func indexHander(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(TEMPATE))

	value, err := loadMD()
	if err != nil {
		panic(err)
	}

	m := map[string]string{
		"Erd": value,
	}

	tmpl.Execute(w, m)
}

func loadMD() (string, error) {
	f, err := os.Open(MD)
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
