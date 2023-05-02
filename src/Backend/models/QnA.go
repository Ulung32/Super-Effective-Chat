package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type QnA struct {
	ID       primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	Question string `bson:"question" json:"question`
	Answer   string `bson:"answer" json:"answer`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}