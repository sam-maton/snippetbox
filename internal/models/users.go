package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

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
	return 0, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
