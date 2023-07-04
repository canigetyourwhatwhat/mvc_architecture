package controllers

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	entity "mvc_go/app/models"
	"net/http"
	"time"
)

func (server *Server) GetInProgressCart(c echo.Context) error {
	sessionId := c.Param("session")
	if sessionId == "" {
		return c.JSON(http.StatusBadRequest, "session is missing")
	}

	session := &entity.Session{ID: sessionId}
	session, err := session.GetSession(server.DB)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Session is not valid")
	}

	if session.ExpiresAt.Before(time.Now()) {
		return c.JSON(http.StatusBadRequest, "Session is already expired")
	}

	var cart *entity.Cart
	cart, err = cart.GetInProgressCartByUserId(server.DB, session.UserID)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, "cart doesn't exist")
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed GetUserInfoByUsername: "+err.Error())
	}

	// collect all the cart items in the cart
	var cartItem entity.CartItem
	cartItems, err := cartItem.GetCarItemsByCartId(server.DB, cart.ID)
	cart.CartItems = cartItems

	return c.JSON(http.StatusOK, cart)
}

func (server *Server) UpdateCart(c echo.Context) error {
	var body entity.UpdateCartItem
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

	var cart *entity.Cart
	cart, err = cart.GetInProgressCartByUserId(server.DB, session.UserID)
	if err != nil && err != sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, "failed GetUserInfoByUsername: "+err.Error())
	}

	// If cart is not created
	if cart == nil {
		err = cart.CreateCart(server.DB, session.UserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to create a cart: "+err.Error())
		}

		cart, err = cart.GetInProgressCartByUserId(server.DB, session.UserID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to get a cart: "+err.Error())
		}
	}

	// reset the cart info
	cart.NetPrice = 0
	cart.TaxPrice = 0
	cart.TotalPrice = 0

	var p entity.Product
	var products []*entity.Product
	var cartItem entity.CartItem
	for _, productInfo := range body.Record {
		product, err := p.GetProductByCode(server.DB, productInfo.ProductCode)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("ProductCode %s doesn't exist", productInfo.ProductCode))
		}
		cartItem.CartId = cart.ID
		cartItem.Quantity = productInfo.Quantity
		cartItem.ProductCode = productInfo.ProductCode
		cartItem.NetPrice = product.Price * float32(productInfo.Quantity)
		cartItem.TaxPrice = product.Price * float32(productInfo.Quantity) * entity.GetTaxPercent()
		cartItem.TotalPrice = product.Price * float32(productInfo.Quantity) * (1 + entity.GetTaxPercent())

		err = cartItem.CreateItemInCart(server.DB)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("ProductCode %s doesn't exist", productInfo.ProductCode))
		}
		products = append(products, product)
		cart.NetPrice += cartItem.NetPrice
		cart.TaxPrice += cartItem.TaxPrice
		cart.TotalPrice += cartItem.TotalPrice
	}

	err = cart.UpdateCart(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update items in the cart: "+err.Error())
	}

	return c.JSON(http.StatusOK, "Updated the cart")
}
