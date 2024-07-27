package user

import (
	"time"
	"vcd-rental/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(username, password string) error
	LoginUser(username, password string) (bool, error)
	GenerateToken(username string) (string, error)
	ValidateToken(tokenStr string) (*jwt.Token, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) RegisterUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := User{
		Username: username,
		Password: string(hashedPassword),
	}

	return s.repo.CreateUser(user)
}

func (s *service) LoginUser(username, password string) (bool, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	return token.SignedString([]byte(config.JWTSecret))
}

func (s *service) ValidateToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.JWTSecret), nil
	})
	return token, err
}
