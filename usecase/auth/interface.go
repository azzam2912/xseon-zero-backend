package auth

import (
	"xseon-zero/domain/model"
)

type AuthUseCase interface {
	Register(email, password string) (*model.User, error)
	Login(email, password string) (string, error) // Returns JWT token
	ValidateToken(token string) (*model.User, error)
}
