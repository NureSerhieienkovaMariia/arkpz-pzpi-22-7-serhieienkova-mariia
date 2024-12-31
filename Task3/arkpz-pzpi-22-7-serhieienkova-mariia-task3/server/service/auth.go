package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"

	"clinic/server/repository"
	"clinic/server/structures"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type TokenClaims struct {
	jwt.StandardClaims
	Id int `json:"id"`
}

type AuthService struct {
	repo repository.AuthorizationRepo
}

func NewAuthService(repo repository.AuthorizationRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user structures.User) (int, error) {
	err := user.Validate()
	if err != nil {
		return 0, fmt.Errorf("user validation error: %w", err)
	}
	user.PasswordHash = generatePasswordHash(user.PasswordHash)
	user.PremiumExpirationDate = "0001-01-01T00:00:00Z"
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (structures.UserToken, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	fmt.Println(fmt.Sprintf("received user: %v", user))
	if err != nil {
		return structures.UserToken{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id: user.Id,
	})

	signedToken, _ := token.SignedString([]byte(signingKey))

	return structures.UserToken{
		Token:  signedToken,
		UserId: user.Id,
	}, nil
}

func (s *AuthService) GetUserById(userId int) (structures.User, error) {
	return s.repo.GetUserById(userId)
}

func (s *AuthService) GenerateTokenByUserId(userId int) (structures.UserToken, error) {
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return structures.UserToken{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id: user.Id,
	})

	signedToken, _ := token.SignedString([]byte(signingKey))

	return structures.UserToken{
		Token:  signedToken,
		UserId: user.Id,
	}, nil
}

func ParseToken(accessToken string) (TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return TokenClaims{}, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return TokenClaims{}, errors.New("token claims are not of type *TokenClaims")
	}

	return *claims, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (structures.UserToken, structures.UserToken, error) {
	// Parse the refresh token
	tokenClaims, err := ParseToken(refreshToken)
	if err != nil {
		return structures.UserToken{}, structures.UserToken{}, err
	}

	// Generate a new access token for the user
	newAccessToken, err := s.GenerateTokenByUserId(tokenClaims.Id)
	if err != nil {
		return structures.UserToken{}, structures.UserToken{}, err
	}

	// Generate a new refresh token for the user
	newRefreshToken, err := s.GenerateTokenByUserId(tokenClaims.Id)
	if err != nil {
		return structures.UserToken{}, structures.UserToken{}, err
	}

	return newAccessToken, newRefreshToken, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
