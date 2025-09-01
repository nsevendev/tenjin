package auth

import (
	"errors"
	"fmt"
	"tenjin/back/internal/crm"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	db     *mongo.Database
	jwtKey []byte
}

func NewAuthService(db *mongo.Database, jwtKey string) *AuthService {
	return &AuthService{
		db:     db,
		jwtKey: []byte(jwtKey),
	}
}

func (s *AuthService) CreateToken(user *crm.User, loginDto LoginDto) (string, error) {
	if user == nil {
		return "", errors.New("identifiants invalides")
	}

	if !user.CheckPassword(loginDto.Password) {
		return "", errors.New("identifiants invalides")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return token, nil
}

func (s *AuthService) NameCookie() string {
	return "1209_!&@_entok_!&@_9834"
}

func (s *AuthService) generateToken(user *crm.User) (string, error) {
	if string(s.jwtKey) == "" {
		return "", errors.New("secret JWT manquant")
	}

	expirationTime := time.Now().Add(24 * time.Hour) // changer ici pour la durée d'expiration

	claims := &tokenClaims{
		IdUser: user.ID.Hex(),
		Email:  user.Email,
		Roles:  user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return tokenString, nil
}
