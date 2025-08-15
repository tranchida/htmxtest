package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"

	"htmxtest/internal/handlers"
	"htmxtest/internal/middleware"
	"htmxtest/internal/router"

	_ "github.com/joho/godotenv/autoload"
)

//go:embed static templates
var staticFiles embed.FS

func main() {

	mux := http.NewServeMux()

	templsParsed, err := template.ParseFS(staticFiles, "templates/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	handlers.Templs = templsParsed

	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	router.Setup(mux, staticFS)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, middleware.LoggingMiddleware(mux)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
