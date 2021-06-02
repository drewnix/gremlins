package keeper

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}



func RegisterHandlers(e *echo.Echo) {
	//e.File("/", "public/index.html")
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	e.GET("/", IndexHandler)
	e.Static("/static", "assets")
	e.GET("/gremlins", GetGremlins())
}

func IndexHandler(c echo.Context) error {
	gremlins := GetGremlins()
	return c.Render(http.StatusOK, "public/index.html", map[string]interface{}{
		"gremlins": gremlins,
	})
}