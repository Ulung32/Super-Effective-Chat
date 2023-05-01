package models

type User struct {
	ID       int64  `bson:"id,omitempty" json:"id,omitempty"`
	UserName string `bson:"username" json:"username`
	Password string `bson:"password" json:"password`
}