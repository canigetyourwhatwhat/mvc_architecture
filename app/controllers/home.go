package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/unrolled/render"
	"net/http"
)

func (server *Server) Home(c echo.Context) error {
	r := render.New(render.Options{
		Layout: "layout",
	})

	_ = r.HTML(c.Response(), http.StatusOK, "home", map[string]interface{}{
		"title": "Home Title",
		"body":  "Home Description",
	})

	return nil
}
