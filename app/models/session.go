package entity

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/jmoiron/sqlx"
	"io"
	"strings"
	"time"
)

type Session struct {
	ID        string    `db:"id"`
	UserID    string    `db:"userId"`
	ExpiresAt time.Time `db:"expiresAt"`
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
	ON DUPLICATE KEY UPDATE id = :id, expiresAt = :expiresAt;
	`
	_, err := db.NamedExec(query, s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) GetSession(db *sqlx.DB) (*Session, error) {
	var session Session
	err := db.Get(&session, "select * from sessions where id = ?", s.ID)
	if err != nil {
		return nil, err
	}
	return &session, nil
}
