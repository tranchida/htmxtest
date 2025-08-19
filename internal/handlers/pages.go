package handlers

import (
	"html/template"
	"log"
	"net/http"

	"htmxtest/internal/services"

	"github.com/go-analyze/charts"
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

	values := [][]float64{
		{2.0, 4.9, 7.0, 23.2, 25.6, 76.7, 135.6, 162.2, 32.6, 20.0, 6.4, 3.3},
		{2.6, 5.9, 9.0, 26.4, 28.7, 70.7, 175.6, 182.2, 48.7, 18.8, 6.0, 2.3},
	}

	opt := charts.NewBarChartOptionWithData(values)
	opt.Title.Text = "Bar Chart"
	opt.XAxis.Labels = []string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	opt.Legend = charts.LegendOption{
		SeriesNames: []string{
			"Rainfall", "Evaporation",
		},
		Offset:       charts.OffsetRight,
		OverlayChart: charts.Ptr(true),
	}

	p := charts.NewPainter(charts.PainterOptions{
		OutputFormat: charts.ChartOutputSVG,
	})

	err := p.BarChart(opt)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	buf, err := p.Bytes()
	if err != nil {
		panic(err)
	}

	w.Write(buf)

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
