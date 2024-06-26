package database

import (
	"context"
	"log"
	"time"

	"github.com/DumanskyiDima/genesis_test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var COLLECTION = "users"

func GetSubscribedUsers() []models.User {
	client := GetClient()
	UserCollection := GetCollection(client, COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var userList []models.User
	cursor, err := UserCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Println(err)
		}
	}()

	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Println(err)
		}
		if user.Status == "active" {
			userList = append(userList, user)
		}
	}

	return userList
}

func FindUser(email string) *models.User {
	client := GetClient()
	UserCollection := GetCollection(client, COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user *models.User
	filter := bson.D{{Key: "email", Value: email}}
	err := UserCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatalln(err)
	}
	log.Println(user)
	return user
}

func CreateUser(email string, status string) (models.User, error) {
	client := GetClient()
	userCollection := GetCollection(client, COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userToPost := models.User{
		ID:               primitive.NewObjectID(),
		Email:            email,
		Status:           status,
		RegistrationDate: primitive.NewDateTimeFromTime(time.Now()),
	}
	_, err := userCollection.InsertOne(ctx, userToPost)
	return userToPost, err
}
