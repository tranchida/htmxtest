package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"time"

	"htmxtest/internal/handlers"
	"htmxtest/internal/middleware"
	"htmxtest/internal/router"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
)

//go:embed static templates
var staticFiles embed.FS

func main() {

	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()

	mux := http.NewServeMux()

	templsParsed, err := template.ParseFS(staticFiles, "templates/*.gohtml")
	if err != nil {
		logger.Fatal().Err(err)
	}
	handlers.Templs = templsParsed

	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		logger.Fatal().Err(err)
	}

	router.Setup(mux, staticFS)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Info().Msgf("Starting server on port %s", port)
	logger.Info().Msg("Press Ctrl+C to stop the server")
	logger.Info().Msgf("Open http://localhost:" + port + "/")
	if err := http.ListenAndServe(":"+port, middleware.LoggingMiddleware(mux, logger)); err != nil {
		logger.Fatal().Msgf("ListenAndServe: %v", err)
	}
}
