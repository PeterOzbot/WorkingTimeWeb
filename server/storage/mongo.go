package storage

import (
	"log"
	"os"
	"workingtimeweb/server/config"
	"workingtimeweb/server/core"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const collectionName = "workingDays"

// MongoRepository : Repository for accessing the database.
type MongoRepository struct {
	logger  *log.Logger
	session *mgo.Session
}

// Find fetches a WorkingDay from mongo according to the query criteria provided.
func (r MongoRepository) Find(repoID string) (*core.WorkingDay, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	var workingDay core.WorkingDay
	err := coll.Find(bson.M{"repoId": repoID, "userId": workingDay.UserID}).One(&workingDay)
	if err != nil {
		r.logger.Printf("error: %v\n", err)
		return nil, err
	}
	return &workingDay, nil
}

// FindAll fetches all workingDays from the database. YES.. ALL! be careful.
func (r MongoRepository) FindAll(selector map[string]interface{}) ([]*core.WorkingDay, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	var workingDays []*core.WorkingDay
	err := coll.Find(selector).All(&workingDays)
	if err != nil {
		r.logger.Printf("error: %v\n", err)
		return nil, err
	}
	return workingDays, nil
}

// Delete deletes a workingDay from mongo according to the query criteria provided.
func (r MongoRepository) Delete(workingDay *core.WorkingDay) error {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	return coll.Remove(bson.M{"repoId": workingDay.RepoID, "userId": workingDay.UserID})
}

// Update updates an workingDay.
func (r MongoRepository) Update(workingDay *core.WorkingDay) error {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	return coll.Update(bson.M{"repoId": workingDay.RepoID, "userId": workingDay.UserID}, workingDay)
}

// Create workingDays in the database.
func (r MongoRepository) Create(workingDays ...*core.WorkingDay) error {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)

	for _, workingDay := range workingDays {
		_, err := coll.Upsert(bson.M{"repoId": workingDay.RepoID, "userId": workingDay.UserID}, workingDay)
		if err != nil {
			return err
		}
	}

	return nil
}

// Count counts documents for a given collection
func (r MongoRepository) Count() (int, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("").C(collectionName)
	return coll.Count()
}

// NewMongoSession dials mongodb and creates a session.
func newMongoSession() (*mgo.Session, error) {
	mongoURL := config.Configuration().MongoURL
	if mongoURL == "" {
		log.Fatal("MongoUrl not provided")
	}
	return mgo.Dial(mongoURL)
}

func newMongoRepositoryLogger() *log.Logger {
	return log.New(os.Stdout, "[mongoDB] ", 0)
}

func NewMongoRepository() core.Repository {
	logger := newMongoRepositoryLogger()
	session, err := newMongoSession()
	if err != nil {
		logger.Fatalf("Could not connect to the database: %v\n", err)
	}

	return MongoRepository{
		session: session,
		logger:  logger,
	}
}
