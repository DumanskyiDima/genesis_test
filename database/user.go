package database

import (
	"context"
	"log"
	"time"

	. "github.com/DumanskyiDima/genesis_test/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var COLLECTION = "users"

func List_Users() []User {
	client := GetClient()
	UserCollection := GetCollection(client, COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var userList []User
	cursor, err := UserCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Fatalln(err)
		}
	}()

	for cursor.Next(ctx) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatalln(err)
		}
		userList = append(userList, user)
	}

	return userList
}

func Find_User(email string) *User {
	client := GetClient()
	UserCollection := GetCollection(client, COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user *User
	filter := bson.D{{Key: "Email", Value: email}}
	err := UserCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil
	}
	return user
}

func Create_User(user User) string {
	client := GetClient()
	userCollection := GetCollection(client, COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userToPost := User{
		ID:     primitive.NewObjectID(),
		Email:  user.Email,
		Status: user.Status,
	}
	result, err := userCollection.InsertOne(ctx, userToPost)
	if err != nil {
		return ""
	}
	return result.InsertedID.(primitive.ObjectID).Hex()
}
