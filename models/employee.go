package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Role     string             `json:"role,omitempty" bson:"role,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone    string             `json:"phone,omitempty" bson:"phone,omitempty"`
}
