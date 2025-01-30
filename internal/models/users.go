package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserModelInterface interface {
	Insert(string, string, string) error
	Authenticate(string, string) (int, error)
	Exists(int) (bool, error)
	Details(int) (User, error)
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	created := time.Now()

	stmt := `INSERT INTO users (name, email, hashed_password, created)
	Values($1, $2, $3, $4)`

	_, err = m.DB.Exec(stmt, name, email, hashedPassword, created)

	if err != nil {
		var postgresError *pq.Error

		if errors.As(err, &postgresError) {
			if postgresError.Code == "23505" && strings.Contains(postgresError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := "SELECT id, hashed_password FROM users WHERE email = $1"

	err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	var exists bool
	stmt := "SELECT EXISTS(SELECT true FROM users WHERE id = $1)"

	err := m.DB.QueryRow(stmt, id).Scan(&exists)

	return exists, err
}

func (m *UserModel) Details(id int) (User, error) {
	var user User
	stmt := "SELECT id, name, email, created FROM users WHERE id = $1"

	err := m.DB.QueryRow(stmt, id).Scan(&user.ID, &user.Name, &user.Email, &user.Created)
	if err != nil {
		if errors.Is(err, ErrNoRecord) {
			return User{}, ErrNoRecord
		} else {
			return User{}, err
		}
	}

	return user, nil
}
