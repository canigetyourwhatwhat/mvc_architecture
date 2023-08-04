package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	entity "mvc_go/app/models"
	"net/http"
)

func (server *Server) AddItemToCart(c echo.Context) error {
	var body entity.CartItemRequest
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
	}

	var session entity.Session
	err := session.ValidateSession(c, server.DB)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var cart *entity.Cart
	cart, err = cart.GetInProgressCartByUserId(server.DB, session.UserID)
	if err != nil && err != sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, "failed GetUserInfoByUsername: "+err.Error())
	}

	// Create cart if doesn't exist
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

	var cartItem *entity.CartItem
	cartItem, err = cartItem.GetCartItemByProductIdAndCartId(server.DB, body.ProductCode, cart.ID)
	if err != nil && err != sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, "failed GetCartItemByProductIdAndCartId: "+err.Error())
	}

	// get price from product
	product := entity.Product{Code: body.ProductCode}
	err = product.GetProductByCode(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to get product: "+err.Error())
	}

	// if already this product is added
	if cartItem != nil {
		cart.NetPrice = cart.NetPrice - cartItem.NetPrice + product.Price*float32(body.Quantity)
		cart.TaxPrice = cart.TaxPrice - cartItem.TaxPrice + product.Price*float32(body.Quantity)*entity.GetTaxPercent()
		cart.TotalPrice = cart.TotalPrice - cartItem.TotalPrice + product.Price*float32(body.Quantity)*(1+entity.GetTaxPercent())

		err = cartItem.DeleteCartItemByCartIdAndCode(server.DB)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed DeleteItemInCart: "+err.Error())
		}
	} else {
		cart.NetPrice = cart.NetPrice + product.Price*float32(body.Quantity)
		cart.TaxPrice = cart.TaxPrice + product.Price*float32(body.Quantity)*entity.GetTaxPercent()
		cart.TotalPrice = cart.TotalPrice + product.Price*float32(body.Quantity)*(1+entity.GetTaxPercent())
	}

	err = cart.UpdateCart(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update item in the cart: "+err.Error())
	}

	cartItem = &entity.CartItem{
		ProductCode: product.Code,
		CartId:      cart.ID,
		Quantity:    body.Quantity,
		NetPrice:    cart.NetPrice,
		TaxPrice:    cart.TaxPrice,
		TotalPrice:  cart.TotalPrice,
		Product:     nil,
	}

	err = cartItem.CreateItemInCart(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create item(s) in a cart: "+err.Error())
	}

	return c.JSON(http.StatusOK, "added product in the cart")
}

func (server *Server) RemoveItemFromCart(c echo.Context) error {
	code := c.Param("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, "product code is missing")
	}

	var session entity.Session
	err := session.ValidateSession(c, server.DB)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
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
	cartItem, err = cartItem.GetCartItemByProductIdAndCartId(server.DB, code, cart.ID)
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

	err = cartItem.DeleteCartItemByCartIdAndCode(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete cart item: "+err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted the product from the cart")
}
