package controllers

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	entity "mvc_go/app/models"
	"net/http"
)

func (server *Server) GetInProgressCart(c echo.Context) error {
	var session entity.Session
	err := session.ValidateSession(c, server.DB)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var cart *entity.Cart
	cart, err = cart.GetInProgressCartByUserId(server.DB, session.UserID)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusOK, "cart doesn't exist")
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed GetUserInfoByUsername: "+err.Error())
	}

	// collect all the cart items in the cart
	cartItem := entity.CartItem{CartId: cart.ID}
	cartItems, err := cartItem.GetCartItemsByCartId(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed GetCartItemsByCartId: "+err.Error())
	}
	cart.CartItems = cartItems

	return c.JSON(http.StatusOK, *cart)
}

func (server *Server) UpdateCart(c echo.Context) error {
	var body entity.UpdateCartItemRequest
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

	// Delete all the items in the cart
	cartItem := entity.CartItem{CartId: cart.ID}
	err = cartItem.DeleteItemInCart(server.DB)

	for _, productInfo := range body.Records {
		product := &entity.Product{Code: productInfo.ProductCode}
		err = product.GetProductByCode(server.DB)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("ProductCode %s doesn't exist", productInfo.ProductCode))
		}

		cartItem = entity.CartItem{
			CartId:      cart.ID,
			Quantity:    productInfo.Quantity,
			ProductCode: productInfo.ProductCode,
			NetPrice:    product.Price * float32(productInfo.Quantity),
			TaxPrice:    product.Price * float32(productInfo.Quantity) * entity.GetTaxPercent(),
			TotalPrice:  product.Price * float32(productInfo.Quantity) * (1 + entity.GetTaxPercent()),
		}

		err = cartItem.CreateItemInCart(server.DB)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed CreateItemInCart: "+err.Error())
		}

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
