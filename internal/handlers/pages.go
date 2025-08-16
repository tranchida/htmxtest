package handlers

import (
	"bytes"
	"html/template"
	"log"
	"math/rand"
	"net/http"

	"htmxtest/internal/services"

	"github.com/wcharczuk/go-chart/v2"
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

func Graph(w http.ResponseWriter, r *http.Request) {
	graph := chart.BarChart{
		Title: "Exemple de graphique",
		Bars: []chart.Value{
			{Value: rand.Float64(), Label: "A"},
			{Value: rand.Float64(), Label: "B"},
			{Value: rand.Float64(), Label: "C"},
		},
	}

	// Buffer pour l'image
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		http.Error(w, "Erreur lors de la génération du graphique", http.StatusInternalServerError)
		return
	}

	// Définir le type de contenu et envoyer l'image
	w.Header().Set("Content-Type", "image/png")
	w.Write(buffer.Bytes())
}

func handleError(w http.ResponseWriter, err error, msg string, code int) {
	log.Printf("%s: %v", msg, err)
	http.Error(w, msg+": "+err.Error(), code)
}

func Render(w http.ResponseWriter, tmpl string, data any) {
	err := Templs.ExecuteTemplate(w, tmpl+".gohtml", data)
	if err != nil {
		handleError(w, err, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func ShowGraph(w http.ResponseWriter, r *http.Request) {

	randomval := rand.Intn(100)
	Render(w, "graph", randomval)
}
