package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"htmxtest/internal/services"
	"htmxtest/internal/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

//go:embed static
var staticFiles embed.FS

func main() {
	r := gin.Default()

	// Servir les fichiers statiques embarqués correctement
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}
	r.StaticFS("/static", http.FS(staticFS))

	r.GET("/", pageHandler)
	r.GET("/about", pageHandler)
	r.GET("/admin", pageHandler)
	r.GET("/randommessage", randomMessageHandler)

	// Gestion du 404 : redirige vers la page d'accueil
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// pageHandler sert la page HTML complète ou juste le contenu pour les requêtes HTMX.
func pageHandler(c *gin.Context) {
	var template templ.Component

	path := c.Request.URL.Path
	switch path {
	case "/":
		template = templates.Home()
	case "/about":
		template = templates.About()
	case "/admin":
		template = templates.Admin()
	default:
		template = templates.Home()
	}

	if c.GetHeader("HX-Request") == "true" {
		Render(c, http.StatusOK, template)
	} else {
		Render(c, http.StatusOK, templates.Layout(template))
	}
}

func Render(ctx *gin.Context, statusCode int, t templ.Component) {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request.Context(), buf); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Data(statusCode, "text/html; charset=utf-8", buf.Bytes())
}

func randomMessageHandler(c *gin.Context) {
	message := services.GetRandomMessage()
	Render(c, http.StatusOK, templates.Message(message))
}
