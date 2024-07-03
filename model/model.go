package model

type User struct {
	Id       string `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Chats struct {
	IdChats   string `json:"id_chats" bson:"idChats"`
	Message   string `json:"message" bson:"message"`
	Responses string `json:"responses" bson:"responses"`
}
