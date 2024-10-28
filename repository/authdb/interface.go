package authdb

import (
	"xseon-zero/domain/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthDBInterface interface {
	BeginTransaction() *gorm.DB
	CreateUser(tx *gorm.DB, user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id uuid.UUID) (*model.User, error)
}
