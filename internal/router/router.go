package router

import (
	"io/fs"
	"net/http"

	"htmxtest/internal/handlers"
)

func Setup(mux *http.ServeMux, staticFS fs.FS) {
	// Static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	// Page routes
	setupPageRoutes(mux)

	// API routes
	setupAPIRoutes(mux)
}

func setupPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.PageHandler)
	mux.HandleFunc("/about", handlers.PageHandler)
	mux.HandleFunc("/admin", handlers.PageHandler)
}

func setupAPIRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/randommessage", handlers.RandomMessageHandler)
	mux.HandleFunc("/showgraph", handlers.ShowGraph)
	mux.HandleFunc("/graph", handlers.Graph)
}
