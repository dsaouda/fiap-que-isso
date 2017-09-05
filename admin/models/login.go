package models

import "gopkg.in/mgo.v2/bson"

type Login struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}

func NewLogin(email string, password string) Login {
	return Login{Email:email, Password: password}
}