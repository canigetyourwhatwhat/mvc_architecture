package entity

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        int       `db:"id"`
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

func (u *User) CreateUser(db *sqlx.DB) error {
	query := `
	INSERT INTO users
	(firstName, lastName, username, password)
	VALUES
	(
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

func (u *User) GetUserInfoById(db *sqlx.DB, id int) error {
	err := db.Get(u, "select * from users where id = ?", id)
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
