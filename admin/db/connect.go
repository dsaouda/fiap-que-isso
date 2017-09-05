package db

import (
	"os"
	"gopkg.in/mgo.v2"
)

var (
	// session mongo
	Session *mgo.Session

	// informacoes do mongo
	Mongo *mgo.DialInfo
)

const (
	//conexão padrão mongo db
	Url = "mongodb://192.168.200.30:27017/fiap_que_isso"
)

func url() string {
	url := os.Getenv("MONGODB_URI")

	if len(url) == 0 {
		url = Url
	}

	return url
}

func Connect() {
	url := url()
	session, err := mgo.Dial(url)

	if err != nil {
		panic(err.Error())
	}

	session.SetSafe(&mgo.Safe{})

	Session = session
	Mongo, _ = mgo.ParseURL(url)
}
