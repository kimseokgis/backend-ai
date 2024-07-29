package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chats struct {
	IdChats   string  `json:"id_chats" bson:"idChats"`
	Message   string  `json:"message" bson:"message"`
	Responses string  `json:"responses" bson:"responses"`
	Score     float64 `json:"score" bson:"score"`
}

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" `
	Username     string             `json:"username" bson:"username"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	Password     string             `json:"password" bson:"password"`
	PasswordHash string             `json:"passwordhash" bson:"passwordhash"`
	Token        string             `json:"token,omitempty" bson:"token,omitempty"`
	Private      string             `json:"private,omitempty" bson:"private,omitempty"`
	Public       string             `json:"public,omitempty" bson:"public,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Response struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type Payload struct {
	User string `json:"user"`
}

type Datasets struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Question string             `json:"question" bson:"questions"`
	Answer   string             `json:"answer" bson:"answer"`
}

type Requests struct {
	Messages string `json:"messages" bson:"messages"`
}

type Secrets struct {
	SecretToken string `json:"secret_token" bson:"secret_token"`
}

type Comment struct {
	CommenterName string `json:"commenter_name" bson:"commenter_name"`
	CommentValue  string `json:"comment_value" bson:"comment_value"`
}
