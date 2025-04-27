package data

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"task_manager/models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func RegisterUser(user models.User) error {
	findOption := bson.D{{Key: "username", Value: user.Username}}
	var existingUser models.User
	err := UserCollection.FindOne(context.TODO(), findOption).Decode(&existingUser)
	if err == nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	user.ID = generateUserID()
	_, err = UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func AuthenticateUser(username, password string) (string, error) {
	findOption := bson.M{"username": username}
	var user models.User
	err := UserCollection.FindOne(context.TODO(), findOption).Decode(&user)
	if err != nil {
		return "", errors.New("doesn't find user")
	}

	if user.Username == username {
		password = strings.TrimSpace(password)
		isPasswordCorrect, err := VerifyPassword(password, user.Password)

		if err != nil {
			fmt.Println(err.Error())

			return "", err
		}
		if !isPasswordCorrect {

			return "", errors.New("incorrect password")
		}
		expirationTime := time.Now().Add(1 * time.Hour)
		claims := &Claims{
			Username: user.Username,
			Role:     user.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Local().Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {

			return "", err
		}

		return tokenString, nil
	}
	return "", errors.New("invalid credentials")
}

// Utility functions
func HashPassword(password string) (string, error) {
	originalPassword := strings.TrimSpace(password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(originalPassword), bcrypt.DefaultCost)
	return string(bytes), err
}

func VerifyPassword(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func generateUserID() string {
	return "user_" + time.Now().Format("20060102150405")
}
