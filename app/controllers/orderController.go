package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	entity "mvc_go/app/models"
	"net/http"
	"strconv"
)

func (server *Server) GetOrder(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "order id is missing")
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong order id")
	}

	var session entity.Session
	err = session.ValidateSession(c, server.DB)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	order := &entity.Order{
		ID: intId,
	}
	err = order.GetOrder(server.DB, intId)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusBadRequest, "Order doesn't exist")
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get order"+err.Error())
	}

	var payment entity.Payment
	err = payment.GetPaymentById(server.DB, order.PaymentId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get payment")
	}

	var cart entity.Cart
	err = cart.GetCartById(server.DB, order.CartId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get payment"+err.Error())
	}

	// collect all the cart items in the cart
	cartItem := entity.CartItem{CartId: cart.ID}
	cartItems, err := cartItem.GetCartItemsByCartId(server.DB)
	cart.CartItems = cartItems

	res := entity.OrderResponse{
		Payment: payment,
		Cart:    cart,
	}

	return c.JSON(http.StatusOK, res)
}

func (server *Server) ListOrders(c echo.Context) error {

	var session entity.Session
	err := session.ValidateSession(c, server.DB)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var cart entity.Cart
	carts, err := cart.GetCartByUserIdAndCompleted(server.DB, session.UserID)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusOK, "No order has been created")
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get carts by user Id"+err.Error())
	}

	var cartItems []entity.CartItem
	for i := range carts {
		cartItem := entity.CartItem{CartId: carts[i].ID}
		newCartItems, err := cartItem.GetCartItemsByCartId(server.DB)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed GetCartItemsByCartId")
		}
		for i := range newCartItems {
			p := entity.Product{Code: newCartItems[i].ProductCode}
			err = p.GetProductByCode(server.DB)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, "Failed GetProductByCode")
			}
			newCartItems[i].Product = &p
		}
		cartItems = append(cartItems, newCartItems...)
	}

	return c.JSON(http.StatusOK, cartItems)
}
