package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	todo "go-todo"
	"go-todo/pkg/repository"
	"time"
)

const salt = "123efe2gfjeklfgjd3"
const expressionTime = 60 * 60 * 24 * 2

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	passwordHash := generatePasswordHash(password)
	user, err := s.repo.GetUser(username, passwordHash)
	if err != nil {
		return "", err
	}

	claims := customClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Unix() + expressionTime,
			Issuer:    "nameOfWebsiteHere",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(salt))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *AuthService) CheckToken(jwtToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(salt), nil
		},
	)

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return "", errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return "", errors.New("jwt is expired")
	}
	return claims.Username, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
