package repository

import (
	mgo "gopkg.in/mgo.v2"
)

type RepositoryService interface {
	SetCollection(collection *mgo.Collection) bool

	Create(docs ...RepositoryService) (err error)
	Read(query RepositoryService) (resultQuery *mgo.Query)
	Update(c *mgo.Collection, docs ...RepositoryService) (err error)
	Delete(c *mgo.Collection, docs ...RepositoryService) (err error)

	/*ReadId(c *mgo.Collection, ID uint64) *mgo.Query
	UpdateId(c *mgo.Collection, ID uint64) error
	DeleteId(c *mgo.Collection, ID uint64) error

	UpdateAll(c *mgo.Collection, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	DeleteAll(c *mgo.Collection, selector interface{}) (info *mgo.ChangeInfo, err error) */
}

type EmployeeService struct {
	Collection *mgo.Collection
}

func (empService *EmployeeService) SetCollection(collection *mgo.Collection) bool {
	if collection != nil {
		empService.Collection = collection
		return true
	}
	return false
}

func (empService *EmployeeService) Create(docs ...RepositoryService) (err error) {
	if empService.Collection != nil {
		err = empService.Collection.Insert(docs)
	}
	return
}

func (empService *EmployeeService) Read(query ...RepositoryService) (resultQuery *mgo.Query) {
	if empService.Collection != nil {
		resultQuery = empService.Collection.Find(query)
	}
	return
}

func (empService *EmployeeService) Update(docs ...RepositoryService) (err error) {
	if empService.Collection != nil {
		err = empService.Collection.Update(docs)
	}
	return
}

func (empService *EmployeeService) Delete(docs ...RepositoryService) (err error) {
	if empService.Collection != nil {
		err = empService.Collection.Remove(docs)
	}
	return
}
