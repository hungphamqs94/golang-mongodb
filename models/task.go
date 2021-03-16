package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	AssignedId int                `json:"assignedId,omitempty" bson:"assignedId,omitempty"`
	Prioty     int                `json:"prioty,omitempty" bson:"prioty,omitempty"`
	Hour       int                `json:"hour,omitempty" bson:"hour,omitempty"`
}
