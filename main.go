package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, LoggingMiddleware(mux)); err != nil {
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
	ctx := r.Context()
	message := services.GetRandomMessage(ctx)
	render(w, "message", message)
}

func handleError(w http.ResponseWriter, err error, msg string, code int) {
	log.Printf("%s: %v", msg, err)
	http.Error(w, msg+": "+err.Error(), code)
}

func render(w http.ResponseWriter, tmpl string, data any) {
	err := templs.ExecuteTemplate(w, tmpl+".go.tmpl", data)
	if err != nil {
		handleError(w, err, "Error executing template", http.StatusInternalServerError)
		return
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	bytes      int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	n, err := lrw.ResponseWriter.Write(b)
	lrw.bytes += n
	return n, err
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ip := r.RemoteAddr
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(lrw, r)
		duration := time.Since(start)
		t := time.Now().Format("02/Jan/2006:15:04:05 -0700")
		log.Printf("%s - - [%s] \"%s %s %s\" %d %d %v",
			ip,
			t,
			r.Method,
			r.URL.RequestURI(),
			r.Proto,
			lrw.statusCode,
			lrw.bytes,
			duration,
		)
	})
}
