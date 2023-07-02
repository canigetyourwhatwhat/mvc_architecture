package controllers

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	entity "mvc_go/app/models"
	"net/http"
	"time"
)

func (server *Server) RegisterUser(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
	}

	_, err := user.GetUserInfoByUsername(server.DB, user.Username)
	if err != nil && err != sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, "Failed GetUserInfoByUsername: "+err.Error())
	}
	if err == nil {
		return c.JSON(http.StatusBadRequest, "This username is already taken ")
	}

	if err = user.HashPassword(); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to hash password: "+err.Error())
	}

	user.ID = uuid.New().String()[:16]

	err = user.CreateUser(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to create user: "+err.Error())
	}

	return c.JSON(http.StatusOK, "User is created")

}

func (server *Server) Login(c echo.Context) error {
	var input entity.LoginInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
	}

	var user *entity.User
	user, err := user.GetUserInfoByUsername(server.DB, input.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "User doesn't exist: ")
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Password is incorrect: "+err.Error())
	}

	var session entity.Session
	sessionID, err := session.CreateSessionID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to create sessionID: "+err.Error())
	}

	session.ID = sessionID
	session.UserID = user.ID
	session.ExpiresAt = time.Now().AddDate(0, 0, 7)

	err = session.CreateOrUpdateSession(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to create session: "+err.Error())
	}

	return c.JSON(http.StatusOK, sessionID)
}
