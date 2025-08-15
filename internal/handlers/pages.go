package handlers

import (
	"html/template"
	"log"
	"net/http"

	"htmxtest/internal/services"
)

var Templs *template.Template

func PageHandler(w http.ResponseWriter, r *http.Request) {

	var tmpl string
	path := r.URL.Path
	switch path {
	case "/":
		tmpl = "home"
	case "/about":
		tmpl = "about"
	case "/admin":
		tmpl = "admin"
	default:
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		Render(w, tmpl, nil)
	} else {
		Render(w, "layout", nil)
	}
}

func RandomMessageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	message := services.GetRandomMessage(ctx)
	Render(w, "message", message)
}

func handleError(w http.ResponseWriter, err error, msg string, code int) {
	log.Printf("%s: %v", msg, err)
	http.Error(w, msg+": "+err.Error(), code)
}

func Render(w http.ResponseWriter, tmpl string, data any) {
	err := Templs.ExecuteTemplate(w, tmpl+".go.tmpl", data)
	if err != nil {
		handleError(w, err, "Error executing template", http.StatusInternalServerError)
		return
	}
}
