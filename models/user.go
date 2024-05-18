package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email            string             `json:"email" bson:"email"`
	RegistrationDate primitive.DateTime `json:"registrationDate" bson:"registrationDate"`
	LastEmailSent    primitive.DateTime `json:"lastEmailSent" bson:"lastEmailSent"`
	Status           string             `json:"status" bson:"status"`
}
