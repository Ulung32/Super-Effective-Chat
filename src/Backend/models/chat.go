package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	ID        	primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	HistoryId 	primitive.ObjectID `bson:"historyid,omitempty" json:"historyid,omitempty"`
	Chat      	string             `bson:"chat" json:"chat"`
	IsBot 		bool `bson:"isbot" json:"isbot"`
}