package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"userId"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"-"`
	Mobile      string             `bson:"mobile,omitempty" json:"mobile,omitempty"`
	DateOfBirth string             `bson:"dateOfBirth,omitempty" json:"dateOfBirth,omitempty"`
}
