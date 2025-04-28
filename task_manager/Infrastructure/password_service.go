package Infrastructure

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, hash string) (bool, error)
	GenerateUserID() string
}
type passwordService struct {
}

func NewPasswordService() PasswordService {
	return &passwordService{}
}

func (p *passwordService) HashPassword(password string) (string, error) {
	originalPassword := strings.TrimSpace(password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(originalPassword), bcrypt.DefaultCost)
	return string(bytes), err
}

func (p *passwordService) VerifyPassword(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *passwordService) GenerateUserID() string {
	return "user_" + time.Now().Format("20060102150405")
}
