package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quote struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Quote      string             `json:"quote,omitempty" bson:"quote,omitempty"`
	AuthorName string             `json:"authorname" bson:"authorname,omitempty"`
}
