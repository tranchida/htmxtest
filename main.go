package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"htmxtest/internal/services"
)

//go:embed static templates
var staticFiles embed.FS

var templs *template.Template

func main() {

	mux := http.NewServeMux()

	templsParsed, err := template.ParseFS(staticFiles, "templates/*.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	templs = templsParsed

	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	mux.HandleFunc("/", pageHandler)

	mux.HandleFunc("/randommessage", randomMessageHandler)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", LoggingMiddleware(mux)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// pageHandler sert la page HTML complète ou juste le contenu pour les requêtes HTMX.
func pageHandler(w http.ResponseWriter, r *http.Request) {

	var template string
	path := r.URL.Path
	switch path {
	case "/":
		template = "home"
	case "/about":
		template = "about"
	case "/admin":
		template = "admin"
	default:
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		render(w, template, nil)
	} else {
		render(w, "layout", nil)
	}
}

func randomMessageHandler(w http.ResponseWriter, r *http.Request) {
	message := services.GetRandomMessage()
	render(w, "message", message)
}

func render(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templs.ExecuteTemplate(w, tmpl + ".go.tmpl", data)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
