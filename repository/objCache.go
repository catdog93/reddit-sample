package repository

import (
	"fmt"
	"github.com/catdog93/GoIT/professions"
	//"log"

	//"gopkg.in/mgo.v2"
	"sync"
)

type ObjCacheEmployeeService interface {
	Add(docs ...professions.EmployeeService) (err error)
	FindId(id uint64) (employee professions.EmployeeService, isCreated bool, err error)
	ReplaceId(employee *professions.EmployeeService) (err error)
	UpsertId(employee *professions.EmployeeService) (err error)
	DeleteId(id uint64) (err error)
	DeleteAll() (err error)
}

type ObjCacheEmployee struct { // Employee.ID equals key of map's element
	Cache map[uint64]professions.EmployeeService
	sync.RWMutex
}

var ObjCache = ObjCacheEmployee{
	Cache: make(map[uint64]professions.EmployeeService),
}

/*
func (c *ObjCacheEmployee) Add(docs ...professions.EmployeeService) (err error) {
	errString := ""
	if docs != nil {
		switch docs.(type) {
		case professions.Employee:

		}
		c.Lock()
		switch {
		case c == nil:
			errString += fmt.Sprintf("cache Add error: CacheEmployee can't be nil\n")
		case c.Cache == nil:
			errString += fmt.Sprintf("cache Add error: cache can't be nil\n")
		default:
			for _, v := range docs {
				if _, ok := c.Cache[v.ID]; ok {
					errString += fmt.Sprintf("cache Add error: existed element with ID = %d\n", v.ID)
				} else {
					c.Cache[v.ID] = v
				}
			}
		}
		c.Unlock()
	} else {
		errString += fmt.Sprintf("cache Add error: impossible to add nil value\n")
	}
	if errString != "" {
		return fmt.Errorf(errString)
	}
	return
}
*/
func (c *ObjCacheEmployee) FindId(id uint64) (employee professions.EmployeeService, isCreated bool, err error) {
	errString := ""
	if id != 0 {
		c.RLock()
		switch {
		case c == nil:
			errString += fmt.Sprintf("cache FindId error: CacheEmployee can't be nil\n")
		case c.Cache == nil:
			errString += fmt.Sprintf("cache FindId error: cache can't be nil\n")
		default:
			if _, ok := c.Cache[id]; ok {
				employee = c.Cache[id]
				isCreated = true
			}
		}
		c.RUnlock()
	} else {
		errString += fmt.Sprintf("cache FindId error: id can't be == 0\n")
	}
	if errString != "" {
		err = fmt.Errorf(errString)
	}
	return
}

/*
func (c *ObjCacheEmployee) ReplaceId(employee *professions.EmployeeService) (err error) {
	errString := ""
	isCreated := false
	if employee != nil {
		c.RLock()
		switch {
		case c == nil:
			c.RUnlock()
			errString += fmt.Sprintf("cache ReplaceId error: CacheEmployee can't be nil\n")
		case c.Cache == nil:
			c.RUnlock()
			errString += fmt.Sprintf("cache ReplaceId error: cache can't be nil\n")
		default:
			c.RUnlock()
			id := employee.ID
			if _, isCreated, err = c.FindId(id); err != nil {
				return
			} else {
				if isCreated {
					c.Lock()
					c.Cache[id] = *employee
					c.Unlock()
				} else {
					errString += fmt.Sprintf("cache ReplaceId error: element doesn't exist with ID = %d\n", id)
				}
			}
		}
	} else {
		errString += fmt.Sprintf("cache ReplaceId error: impossible to replace employee with nil value\n")
	}
	if errString != "" {
		err = fmt.Errorf(errString)
	}
	return
}

func (c *ObjCacheEmployee) UpsertId(employee *professions.EmployeeService) (err error) {
	isCreated := false
	if employee != nil {
		if _, isCreated, err = c.FindId(employee.ID); err != nil {
			return
		} else {
			if isCreated {
				err = c.ReplaceId(employee)
			} else {
				err = c.Add(*employee)
			}
		}
	} else {
		fmt.Errorf("cache UpsertId error: impossible to upsert employee with nil value\n")
	}
	return
}
*/
func (c *ObjCacheEmployee) DeleteId(id uint64) (err error) {
	isCreated := false
	if _, isCreated, err = c.FindId(id); err != nil {
		return
	} else {
		if isCreated {
			c.Lock()
			delete(c.Cache, id)
			c.Unlock()
		} else {
			err = fmt.Errorf("cache DeleteId error: impossible to delete element with id = %v because it doesn't exist\n", id)
		}
	}
	return
}

func (c *ObjCacheEmployee) DeleteAll() (err error) {
	errString := ""
	c.RLock()
	switch {
	case c == nil:
		c.RUnlock()
		errString += fmt.Sprintf("cache DeleteAll error: CacheEmployee can't be nil\n")
	case c.Cache == nil:
		c.RUnlock()
		errString += fmt.Sprintf("cache DeleteAll error: cache can't be nil\n")
	default:
		c.RUnlock()
		c.Lock()
		c.Cache = make(map[uint64]professions.EmployeeService)
		c.Unlock()
	}
	return
}
