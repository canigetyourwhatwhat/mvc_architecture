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

	var cartItem *entity.CartItem
	cartItem, err = cartItem.GetCartItemByProductIdAndCartId(server.DB, body.ProductCode, cart.ID)
	if err != nil && err != sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, "failed GetUserInfoByUsername: "+err.Error())
	}
	if cartItem != nil {
		return c.JSON(http.StatusBadRequest, "This product is already added, please use update to update quantity")
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

	cart.NetPrice = cart.NetPrice + product.Price*float32(body.Quantity)
	cart.TaxPrice = cart.TaxPrice + product.Price*float32(body.Quantity)*entity.GetTaxPercent()
	cart.TotalPrice = cart.TotalPrice + product.Price*float32(body.Quantity)*(1+entity.GetTaxPercent())

	cartItem = &entity.CartItem{
		ProductCode: product.Code,
		CartId:      cart.ID,
		Quantity:    body.Quantity,
		NetPrice:    cart.NetPrice,
		TaxPrice:    cart.TaxPrice,
		TotalPrice:  cart.TotalPrice,
	}
	err = cartItem.CreateItemInCart(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create item(s) in a cart: "+err.Error())
	}

	err = cart.UpdateCart(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update item in the cart: "+err.Error())
	}

	return c.JSON(http.StatusOK, "added product in the cart")
}

func (server *Server) RemoveItemFromCart(c echo.Context) error {
	var body entity.DeleteCartItemRequest
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
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, "cart doesn't exist")
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed GetUserInfoByUsername: "+err.Error())
	}

	// calculate price
	var cartItem *entity.CartItem
	cartItem, err = cartItem.GetCartItemByProductIdAndCartId(server.DB, body.ProductCode, cart.ID)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, "This product is not in the cart")
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to get product: "+err.Error())
	}

	cart.NetPrice = cart.NetPrice - cartItem.NetPrice
	cart.TaxPrice = cart.TaxPrice - cartItem.TaxPrice
	cart.TotalPrice = cart.TotalPrice - cartItem.TotalPrice

	err = cart.UpdateCart(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update item in the cart: "+err.Error())
	}

	err = cartItem.DeleteItemInCart(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete cart item: "+err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted the product from the cart")
}
