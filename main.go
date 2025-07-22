package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"htmxtest/internal/services"
	"htmxtest/internal/templates"

	"github.com/a-h/templ"
)

//go:embed static
var staticFiles embed.FS

func main() {

	// Servir les fichiers statiques embarqués correctement
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

    http.HandleFunc("/", pageHandler)

    http.HandleFunc("/randommessage", randomMessageHandler)

    log.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

// pageHandler sert la page HTML complète ou juste le contenu pour les requêtes HTMX.
func pageHandler(w http.ResponseWriter, r *http.Request) {
	var template templ.Component

    path := r.URL.Path
	switch path {
	case "/":
		template = templates.Home()
	case "/about":
		template = templates.About()
	case "/admin":
		template = templates.Admin()
	default:
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

    if r.Header.Get("HX-Request") == "true" {
		template.Render(r.Context(), w)
    } else {
        templates.Layout(template).Render(r.Context(), w)
    }
}

func randomMessageHandler(w http.ResponseWriter, r *http.Request) {
	message := services.GetRandomMessage()
	templates.Message(message).Render(r.Context(), w)
}
