package repository

import (
	mgo "gopkg.in/mgo.v2"
)

type Obj map[string]interface{}

type ID struct {
	ID uint64 `json:"id" bson:"_id"`
}

func Create(collection *mgo.Collection, docs ...interface{}) (err error) {
	if collection != nil {
		err = collection.Insert(docs...)
	}
	return
}

func ReadAll(results []Obj, collection *mgo.Collection) (err error) {
	if collection != nil {
		err = collection.Find(Obj{}).All(&results)
	}
	return
}

func ReadId(ID uint64, result *Obj, collection *mgo.Collection) (err error) {
	if collection != nil && ID > 0 {
		err = collection.FindId(ID).One(result)
	}
	return
}

func UpdateId(ID uint64, update interface{}, collection *mgo.Collection) (err error) {
	if collection != nil && ID > 0 {
		err = collection.UpdateId(ID, update)
	}
	return
}

func Update(selector interface{}, update interface{}, collection *mgo.Collection) (err error) {
	if collection != nil {
		err = collection.Update(selector, update)
	}
	return
}

func DeleteId(ID uint64, collection *mgo.Collection) (err error) {
	if collection != nil && ID > 0 {
		err = collection.RemoveId(ID)
	}
	return
}

func Delete(selector interface{}, collection *mgo.Collection) (err error) {
	if collection != nil {
		err = collection.Remove(selector)
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
