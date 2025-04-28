package Infrastructure

import (
	domain "task_manager/Domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userID string, role string) (string, error)
	ValidateToken(tokenString string, claims *domain.Claims) (bool, error)
}
type jwtService struct {
	secretKey string
}

func NewJWTService(secretKey string) JWTService {
	return &jwtService{
		secretKey: secretKey,
	}
}
func (j *jwtService) GenerateToken(username string, role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &domain.Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {

		return "", err
	}

	return tokenString, nil

}
func (j *jwtService) ValidateToken(tokenString string, claims *domain.Claims) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
