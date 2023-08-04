package entity

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"io"
	"strings"
	"time"
)

type Session struct {
	ID        string    `db:"id"`
	UserID    int       `db:"userId"`
	ExpiresAt time.Time `db:"expiresAt"`
}

func (s *Session) ValidateSession(c echo.Context, db *sqlx.DB) error {
	s.ID = c.Request().Header.Get("session")
	if s.ID == "" {
		return errors.New("session is missing")
	}
	err := s.GetSession(db)
	if err != nil {
		return errors.New("session is not valid")
	}

	if s.ExpiresAt.Before(time.Now()) {
		return errors.New("session is already expired")
	}
	return nil
}

func (s *Session) CreateSessionID() (string, error) {
	sidByte := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, sidByte)
	if err != nil {
		return "", err
	}
	sessionID := strings.TrimRight(base32.StdEncoding.EncodeToString(sidByte), "=")

	return sessionID, nil
}

func (s *Session) CreateOrUpdateSession(db *sqlx.DB) error {
	query := `
	INSERT INTO sessions
	(
	 	id,
		userId,
		expiresAt
	)
	VALUES
	(
	 	:id,
		:userId,
		:expiresAt
	)
	ON DUPLICATE KEY UPDATE id = :id, userId = :userId, expiresAt = :expiresAt;
	`
	_, err := db.NamedExec(query, s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) GetSession(db *sqlx.DB) error {
	err := db.Get(s, "select * from sessions where id = ?", s.ID)
	if err != nil {
		return err
	}
	return nil
}
