package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type History struct {
	ID       primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	UserId   primitive.ObjectID `bson:"userid,omitempty" json:"userid,omitempty"`
	Nama     string             `bson:"name" json:"name"`
}
