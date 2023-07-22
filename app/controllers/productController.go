package controllers

import (
	"github.com/labstack/echo/v4"
	entity "mvc_go/app/models"
	"net/http"
	"strconv"
)

func (server *Server) ListProducts(c echo.Context) error {

	q := c.Request().URL.Query()
	page, err := strconv.Atoi(q.Get("page"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "page number is not incorrect"+err.Error())
	}
	if page <= 0 {
		page = 1
	}

	perPage := 10

	productModel := entity.Product{}
	products, totalRows, err := productModel.GetProducts(server.DB, perPage, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, struct {
		Products []entity.Product
		Count    int
	}{
		Products: products,
		Count:    totalRows,
	})
}

func (server *Server) GetProductByCode(c echo.Context) error {

	code := c.Param("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, "product code is missing")
	}

	productModel := entity.Product{}
	product, err := productModel.GetProductByCode(server.DB, code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	return c.JSON(http.StatusOK, product)
}
