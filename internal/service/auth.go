package service

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"todoApp/internal/model"
	"todoApp/internal/repository"
)

type keys struct {
	jwtKey    string
	salt      string
	expiresAt int
}

const expiresAt = 12

type AuthService struct {
	repo *repository.Repository
	keys
}

func NewAuthService(repo *repository.Repository, jwtKey, passwordSalt string, expiresAt int) *AuthService {
	return &AuthService{
		repo: repo,
		keys: keys{
			jwtKey:    jwtKey,
			salt:      passwordSalt,
			expiresAt: expiresAt,
		},
	}
}

func (s *AuthService) ParseToken(accesToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accesToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign method")
		}
		return []byte(s.jwtKey), nil
	})

	if err != nil {
		return -1, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return -1, errors.New("invalid claims")
	}

	return claims.UserId, nil
}

func (s *AuthService) CreateUser(user model.User) (int64, error) {
	user.Password = s.createHashPassword(user.Password)
	return s.repo.AuthorizationRepository.AddUser(user)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"userId"`
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	u, err := s.repo.AuthorizationRepository.GetUser(username, s.createHashPassword(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiresAt * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: u.Id,
	})

	return token.SignedString([]byte(s.jwtKey))
}

func (s *AuthService) createHashPassword(pass string) string {
	sha := sha512.New()
	sha.Write([]byte(pass))

	return fmt.Sprintf("%x", sha.Sum([]byte(s.salt)))
}
