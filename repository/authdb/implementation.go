package authdb

import (
	"log"
	"xseon-zero/domain/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authDBImpl struct {
	db *gorm.DB
}

func NewAuthDBImpl(db *gorm.DB) AuthDBInterface {
	return &authDBImpl{db: db}
}

func (a *authDBImpl) BeginTransaction() *gorm.DB {
	return a.db.Begin()
}

func (a *authDBImpl) CreateUser(tx *gorm.DB, user *model.User) error {
	result := tx.Create(user)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return result.Error
	}
	return nil
}

func (a *authDBImpl) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := a.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		log.Printf("Error getting user by email: %v", result.Error)
		return nil, result.Error
	}
	return &user, nil
}

func (a *authDBImpl) GetUserByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	result := a.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		log.Printf("Error getting user by ID: %v", result.Error)
		return nil, result.Error
	}
	return &user, nil
}
