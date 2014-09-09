package top

import (
	"html/template"
	"net/http"
)

var rootTemplate *template.Template

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, nil)
}

func init() {
	rootTemplate = template.Must(template.ParseFiles("views/root.json"))
	http.HandleFunc("/", rootHandler)
}
