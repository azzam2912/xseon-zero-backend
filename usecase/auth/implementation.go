package auth

import (
	"errors"
	"log"
	"time"
	"xseon-zero/domain/model"
	"xseon-zero/repository/authdb"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authImpl struct {
	authDB    authdb.AuthDBInterface
	jwtSecret []byte
}

func NewAuthImpl(authDB authdb.AuthDBInterface, jwtSecret string) AuthUseCase {
	return &authImpl{
		authDB:    authDB,
		jwtSecret: []byte(jwtSecret),
	}
}

func (a *authImpl) Register(email, password string) (*model.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, err
	}

	user := &model.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashedPassword),
	}

	tx := a.authDB.BeginTransaction()
	err = a.authDB.CreateUser(tx, user)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
		return nil, err
	}

	return user, nil
}

func (a *authImpl) Login(email, password string) (string, error) {
	user, err := a.authDB.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	tokenString, err := token.SignedString(a.jwtSecret)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return "", err
	}

	return tokenString, nil
}

func (a *authImpl) ValidateToken(tokenString string) (*model.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return a.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			return nil, err
		}

		user, err := a.authDB.GetUserByID(userID)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, errors.New("invalid token")
}
