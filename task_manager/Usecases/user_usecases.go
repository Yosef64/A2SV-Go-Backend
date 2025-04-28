package usecases

import (
	"errors"
	"fmt"
	Domain "task_manager/Domain"
	"task_manager/Infrastructure"
	"task_manager/Repositories"
)

type UserUsecase interface {
	Register(user Domain.User) error
	Login(username, password string) (string, error)
}
type userUsecase struct {
	userRepo        Repositories.UserRepository
	jwtService      Infrastructure.JWTService
	passwordService Infrastructure.PasswordService
}

func NewUserUsecase(userRepo Repositories.UserRepository, jwtService Infrastructure.JWTService, passwordService Infrastructure.PasswordService) UserUsecase {
	return &userUsecase{
		userRepo:        userRepo,
		jwtService:      jwtService,
		passwordService: passwordService,
	}
}
func (u *userUsecase) Register(user Domain.User) error {
	return u.userRepo.Register(user)
}
func (u *userUsecase) Login(username, password string) (string, error) {
	user, err := u.userRepo.FindUserById(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		fmt.Println("User not found")
		return "", errors.New("invalid username or password")
	}

	if isValid, err := u.passwordService.VerifyPassword(user.Password, password); err != nil || !isValid {
		if err != nil {
			fmt.Println("Error verifying password:", err)
		}
		fmt.Println("Invalid password")
		return "", errors.New("invalid username or password")
	}
	token, err := u.jwtService.GenerateToken(user.Username, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
