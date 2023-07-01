package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (server *Server) AddItemToCart(c echo.Context) error {

	return c.JSON(http.StatusOK, "added")
}
