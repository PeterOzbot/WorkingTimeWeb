package main

import (
	"log"
	"workingtimeweb/server/config"

	"github.com/globalsign/mgo"
)

func main() {
	var err error
	mongoURL := config.Configuration().MongoURL
	if mongoURL == "" {
		log.Fatal("MongoUrl not provided")
	}
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	err = session.DB("").AddUser("mongo_user", "mongo_secret", false)

	info := &mgo.CollectionInfo{}
	err = session.DB("").C("kudos").Create(info)

	if err != nil {
		log.Fatal(err)
	}
}
