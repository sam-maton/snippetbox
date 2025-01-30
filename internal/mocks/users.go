package mocks

import (
	"time"

	"github.com/sam-maton/snippetbox/internal/models"
)

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@email.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "test@email.com" && password == "password123" {
		return 1, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Details(id int) (models.User, error) {
	switch id {
	case 1:
		return models.User{
			ID:      1,
			Name:    "Tester",
			Email:   "test@email.com",
			Created: time.Now(),
		}, nil
	default:
		return models.User{}, nil
	}
}
