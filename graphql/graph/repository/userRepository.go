package repository

import (
	"context"
	"log"
	"time"

	"github.com/KinyaElGrande/Go-exercise-101/graphql/graph/database"
	"github.com/KinyaElGrande/Go-exercise-101/graphql/graph/jwt"
	"github.com/KinyaElGrande/Go-exercise-101/graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var usersCollections *mongo.Collection = database.OpenCollection(database.Client, "users")

//HashPassword hashes  a given password
func HashPassword(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashBytes), err
}

//VerifyPassword compares the hashed password with the passed in password
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//RegisterUser creates new users in the database
func RegisterUser(input *model.NewUser) string {
	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		log.Fatal("Error in hashing your password")
	}

	user.Username = input.Username
	user.Password = hashedPassword
	_, err = usersCollections.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		log.Fatal(err)
	}

	return token
}

func GetUserIdByUsername(username string) (string, error) {
	var user model.User
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	documentReturned := usersCollections.FindOne(ctx, bson.M{"username": &username})
	defer cancel()

	documentReturned.Decode(&user)

	return user.ID, nil
}
