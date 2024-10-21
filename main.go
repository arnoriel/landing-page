// main.go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-page/database"
	"go-page/handlers"
	"html/template"
	"io"
)

// Template struct untuk rendering HTML
type Template struct {
	templates *template.Template
}

// Render method untuk Template struct
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Inisialisasi koneksi database
	database.InitDB()

	// Inisialisasi Echo framework
	e := echo.New()

	// Inisialisasi template rendering
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routing untuk web interface
	e.GET("/", handlers.IndexPage)
	e.GET("/faq", handlers.FAQPage)
	e.POST("/faqs", handlers.CreateFAQ)
	e.POST("/faqs/update", handlers.UpdateFAQ)
	e.POST("/faqs/delete", handlers.DeleteFAQ)
	e.GET("/settings", handlers.SettingsPage)
	e.POST("/settings/update", handlers.UpdateSettings)

	// Jalankan server
	e.Logger.Fatal(e.Start(":8080"))
}
