package view

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, paths []string, vars map[string]interface{}) {
	tmpl, err := template.ParseFiles(paths...)
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, vars)
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
	}
}
