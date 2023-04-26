package models

type QnA struct {
	ID       int64  `bson:"id,omitempty" json:"id,omitempty"`
	Question string `bson:"question" json:"question`
	Answer   string `bson:"answer" json:"answer`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}