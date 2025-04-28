package Repositories

import (
	"context"
	"errors"
	Domain "task_manager/Domain"

	"task_manager/Infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Register(user Domain.User) error
	FindUserById(username string) (*Domain.User, error)
}
type userRepository struct {
	collection      *mongo.Collection
	passwordService Infrastructure.PasswordService
}

func NewUserRepository(collection *mongo.Collection, passwordService Infrastructure.PasswordService) UserRepository {
	return &userRepository{collection: collection, passwordService: passwordService}
}
func (r *userRepository) Register(user Domain.User) error {
	findOption := bson.D{{Key: "username", Value: user.Username}}
	var existingUser Domain.User
	err := r.collection.FindOne(context.TODO(), findOption).Decode(&existingUser)
	if err == nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := r.passwordService.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	user.ID = r.passwordService.GenerateUserID()
	_, err = r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindUserById(username string) (*Domain.User, error) {
	filter := bson.M{
		"username": username,
	}
	var user Domain.User
	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
