package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"htmxtest/internal/templates"
	"htmxtest/internal/services"
)

//go:embed static
var staticFiles embed.FS

func main() {

	e := echo.New()
	e.StaticFS("/static", echo.MustSubFS(staticFiles, "static"))
	e.GET("/", pageHandler)
	e.GET("/about", pageHandler)
	e.GET("/admin", pageHandler)
	e.GET("/randommessage", randomMessageHandler)

	e.HTTPErrorHandler = func (err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok && he.Code == http.StatusNotFound {
			c.Redirect(http.StatusFound, "/")
			return
		}
		c.Logger().Error(err)
	}

	if err := e.Start(":8080"); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// pageHandler serves the full HTML page (layout + content) or just the content for HTMX requests.
// It determines which content to show based on the URL path.
func pageHandler(c echo.Context) error {

	var template templ.Component

	path := c.Request().URL.Path
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

	if c.Request().Header.Get("HX-Request") == "true" {
		return Render(c, http.StatusOK, template)
	} else {
		// For full page loads, render the full layout
		return Render(c, http.StatusOK, templates.Layout(template))
	}
}
			

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func randomMessageHandler(c echo.Context) error {
	message := services.GetRandomMessage()
	return Render(c, http.StatusOK, templates.Message(message))
}




