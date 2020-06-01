package repository

import (
	mgo "gopkg.in/mgo.v2"
)

type Obj map[string]interface{}

type ID struct {
	ID uint64 `json:"id" bson:"_id"`
}

// service declares CRUD operations
type Repository interface {
	Connect() error

	Create(docs ...interface{}) (err error)
	Read(query interface{}) (resultQuery *mgo.Query)
	Update(selector interface{}, update interface{}) (err error)
	Delete(selector ...interface{}) (err error)

	ReadId(ID uint64) (resultQuery *mgo.Query)
	UpdateId(ID uint64, update interface{}) (err error)
	DeleteId(ID uint64) (err error)

	UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	DeleteAll(selector interface{}) (info *mgo.ChangeInfo, err error)
}

//func Connect(url, dbName, collectionName string) error {
//	session, err := mgo.Dial(url)
//	if err != nil {
//		return err
//	}
//	defer session.Close()
//	ai.Connect(session.DB(dbName).C(collectionName))
//	return nil
//}

func Create(collection *mgo.Collection, docs ...interface{}) (err error) {
	if collection != nil {
		err = collection.Insert(docs...)
	}
	return
}

func ReadAll(results *[]Obj, collection *mgo.Collection) (err error) {
	if collection != nil {
		err = collection.Find(Obj{}).All(results)
	}
	return
}

func Update(selector interface{}, update interface{}, collection *mgo.Collection) (err error) {
	if collection != nil {
		err = collection.Update(selector, update)
	}
	return
}

func Delete(selector interface{}, collection *mgo.Collection) (err error) {
	if collection != nil {
		err = collection.Remove(selector)
	}
	return
}

func ReadId(ID ID, results []Obj, collection *mgo.Collection) (err error) {
	if collection != nil && ID.ID > 0 {
		err = collection.FindId(ID.ID).All(results)
	}
	return
}

func ReadIdd(ID ID, result *Obj, collection *mgo.Collection) (err error) {
	if collection != nil && ID.ID > 0 {
		err = collection.FindId(ID.ID).One(result)
	}
	return
}

func UpdateId(ID uint64, update interface{}, collection *mgo.Collection) (err error) {
	if collection != nil && ID > 0 {
		err = collection.UpdateId(ID, update)
	}
	return
}

func DeleteId(ID uint64, collection *mgo.Collection) (err error) {
	if collection != nil && ID > 0 {
		err = collection.RemoveId(ID)
	}
	return
}

func DeleteAll(collection *mgo.Collection) (info *mgo.ChangeInfo, err error) {
	if collection != nil {
		info, err = collection.RemoveAll(Obj{})
	}
	return
}

func UpdateAll(selector interface{}, update interface{}, collection *mgo.Collection) (info *mgo.ChangeInfo, err error) {
	if collection != nil {
		info, err = collection.UpdateAll(selector, update)
	}
	return
}
