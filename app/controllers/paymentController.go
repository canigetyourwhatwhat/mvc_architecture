package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	entity "mvc_go/app/models"
	"net/http"
)

func (server *Server) CreatePayment(c echo.Context) error {

	var body entity.MakePaymentInput
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
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusBadRequest, "cart doesn't exist")

	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed GetUserInfoByUsername: "+err.Error())
	}

	payment := &entity.Payment{
		UserId: session.UserID,
		Amount: cart.TotalPrice,
		Method: entity.PaymentMethod(body.PaymentMethod),
	}
	paymentId, err := payment.CreatePayment(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed creating payment: "+err.Error())
	}

	err = payment.GetPaymentById(server.DB, paymentId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed creating payment: "+err.Error())
	}

	var user entity.User
	err = user.GetUserInfoById(server.DB, session.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed getting user: "+err.Error())
	}

	order := entity.Order{
		UserId:    cart.UserId,
		CartId:    cart.ID,
		PaymentId: paymentId,
	}
	err = order.CreateOrder(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed creating order: "+err.Error())
	}

	cart.Status = entity.Completed
	err = cart.UpdateCart(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed updating cart: "+err.Error())
	}

	return c.JSON(http.StatusInternalServerError, "payment has been made")
}
