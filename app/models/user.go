package entity

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"io"
	"strings"
	"time"
)

type User struct {
	ID        string    `db:"id"`
	FirstName string    `db:"firstName" json:"FirstName"`
	LastName  string    `db:"lastName" json:"LastName"`
	Username  string    `db:"username" json:"Username"`
	Password  string    `db:"password" json:"Password"`
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}

type LoginInput struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type Session struct {
	ID        string    `db:"id"`
	UserID    string    `db:"userId"`
	ExpiresAt time.Time `db:"expiresAt"`
}

func (u *User) CreateUser(db *sqlx.DB) error {
	query := `
	INSERT INTO users
	(id, firstName, lastName, username, password)
	VALUES
	(
	 	:id,	
		:firstName,
	 	:lastName,
	 	:username,
	 	:password
	)
	`
	_, err := db.NamedExec(query, *u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) GetUserInfoByUsername(db *sqlx.DB, username string) (*User, error) {
	var user User
	err := db.Get(&user, "select * from users where username = ?", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) CreateSessionID() (string, error) {
	sidByte := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, sidByte)
	if err != nil {
		return "", err
	}
	sessionID := strings.TrimRight(base32.StdEncoding.EncodeToString(sidByte), "=")

	return sessionID, nil
}

func (u *User) CreateSession(db *sqlx.DB, session *Session) error {
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
	`
	_, err := db.NamedExec(query, session)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}
