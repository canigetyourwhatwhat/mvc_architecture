package controllers

import (
	"database/sql"
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

	var emptyCart entity.Cart
	cart, err := emptyCart.GetInProgressCartByUserId(server.DB, session.UserID)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, "cart doesn't exist")
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed GetUserInfoByUsername: "+err.Error())
	}

	return c.JSON(http.StatusOK, cart)
}
