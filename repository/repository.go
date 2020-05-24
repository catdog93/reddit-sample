package repository

import (
	mgo "gopkg.in/mgo.v2"
)

type Obj map[string]interface{}

// service declares CRUD operations
type RepositoryService interface {
	SetCollection(collection *mgo.Collection) bool

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

// struct implements CRUD operations, it can be used for any Collection
type ProfessionsService struct {
	Collection *mgo.Collection
}

func (empService *ProfessionsService) SetCollection(collection *mgo.Collection) bool {
	if collection != nil {
		empService.Collection = collection
		return true
	}
	return false
}

func (empService *ProfessionsService) Create(docs ...interface{}) (err error) {
	if empService.Collection != nil {
		err = empService.Collection.Insert(docs...)
	}
	return
}

func (empService *ProfessionsService) ReadAll(results *[]Obj) (err error) {
	if empService.Collection != nil {
		err = empService.Collection.Find(Obj{}).All(results)
	}
	return
}

func (empService *ProfessionsService) Update(selector interface{}, update interface{}) (err error) {
	if empService.Collection != nil {
		err = empService.Collection.Update(selector, update)
	}
	return
}

func (empService *ProfessionsService) Delete(selector interface{}) (err error) {
	if empService.Collection != nil {
		err = empService.Collection.Remove(selector)
	}
	return
}

func (empService *ProfessionsService) ReadId(ID uint64, results *[]Obj) (err error) {
	if empService.Collection != nil && ID > 0 {
		err = empService.Collection.FindId(ID).All(results)
	}
	return
}

func (empService *ProfessionsService) UpdateId(ID uint64, update interface{}) (err error) {
	if empService.Collection != nil && ID > 0 {
		err = empService.Collection.UpdateId(ID, update)
	}
	return
}

func (empService *ProfessionsService) DeleteId(ID uint64) (err error) {
	if empService.Collection != nil && ID > 0 {
		err = empService.Collection.RemoveId(ID)
	}
	return
}

func (empService *ProfessionsService) DeleteAll() (info *mgo.ChangeInfo, err error) {
	if empService.Collection != nil {
		info, err = empService.Collection.RemoveAll(Obj{})
	}
	return
}

func (empService *ProfessionsService) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	if empService.Collection != nil {
		info, err = empService.Collection.UpdateAll(selector, update)
	}
	return
}
