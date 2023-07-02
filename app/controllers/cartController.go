package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	entity "mvc_go/app/models"
	"net/http"
	"time"
)

func (server *Server) AddItemToCart(c echo.Context) error {
	var body entity.AddCartItemRequest
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
	}

	session := &entity.Session{ID: body.SessionId}
	session, err := session.GetSession(server.DB)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Session is not valid")
	}

	if session.ExpiresAt.Before(time.Now()) {
		return c.JSON(http.StatusBadRequest, "Session is already expired")
	}

	var emptyCart entity.Cart
	cart, err := emptyCart.GetInProgressCartByUserId(server.DB, session.UserID)
	if err != nil && err != sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, "failed GetUserInfoByUsername: "+err.Error())
	}

	// If cart is not created
	if cart == nil {
		err = emptyCart.CreateCart(server.DB, session.UserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to create a cart: "+err.Error())
		}

		cart, err = cart.GetInProgressCartByUserId(server.DB, session.UserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to get a cart: "+err.Error())
		}
	}

	// calculate price
	var product *entity.Product
	product, err = product.GetProductByCode(server.DB, body.ProductCode)
	if err != nil && err != sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, "failed to get product: "+err.Error())
	}

	NetPrice := cart.NetPrice + product.Price*float32(body.Quantity)
	TaxPrice := cart.TaxPrice + product.Price*float32(body.Quantity)*entity.GetTaxPercent()
	TotalPrice := cart.TotalPrice + product.Price*float32(body.Quantity)*(1+entity.GetTaxPercent())

	cartItem := &entity.CartItem{
		ProductId:  product.ID,
		CartId:     cart.ID,
		Quantity:   body.Quantity,
		NetPrice:   NetPrice,
		TaxPrice:   TaxPrice,
		TotalPrice: TotalPrice,
	}
	err = cartItem.CreateItemInCart(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create item(s) in a cart: "+err.Error())
	}

	return c.JSON(http.StatusOK, "added product in the cart")
}
