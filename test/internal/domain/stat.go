package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Stat struct {
	Ts primitive.DateTime `json:"ts" bson:"ts"`
	V  int                `json:"v" bson:"v"`
}
