package repository

import (
	"fmt"
	"github.com/catdog93/GoIT/professions"
	//"log"

	//"gopkg.in/mgo.v2"
	"sync"
)

type CacheEmployee struct { // Employee.ID equals key of map's element
	Cache map[uint64]professions.Employee
	sync.RWMutex
}

var Cache = CacheEmployee{
	Cache: make(map[uint64]professions.Employee),
}

func (c *CacheEmployee) Add(docs ...professions.Employee) (err error) {
	errString := ""
	switch {
	case c == nil:
		errString += fmt.Sprintf("cache Add error: CacheEmployee can't be nil\n")
	case c.Cache == nil:
		errString += fmt.Sprintf("cache Add error: cache can't be nil\n")
	case docs == nil:
		errString += fmt.Sprintf("cache Add error: impossible to add nil value\n")
	default:
		c.Lock()
		for _, v := range docs {
			if _, ok := c.Cache[v.ID]; ok {
				errString += fmt.Sprintf("cache Add error: existed element with ID = %d\n", v.ID)
			} else {
				c.Cache[v.ID] = v
			}
		}
		c.Unlock()
		if errString != "" {
			return fmt.Errorf(errString)
		}
	}
	return
}

func (c *CacheEmployee) FindId(id uint64) (employee professions.Employee, isCreated bool, err error) {
	errString := ""
	c.RLock()
	switch {
	case c == nil:
		errString += fmt.Sprintf("cache FindId error: CacheEmployee can't be nil\n")
	case c.Cache == nil:
		errString += fmt.Sprintf("cache FindId error: cache can't be nil\n")
	case id == 0:
		errString += fmt.Sprintf("cache FindId error: id can't be == 0\n")
	default:
		if _, ok := c.Cache[id]; ok {
			employee = c.Cache[id]
			isCreated = true
		}
	}
	c.RUnlock()
	if errString != "" {
		err = fmt.Errorf(errString)
	}
	return
}

func (c *CacheEmployee) ReplaceId(employee *professions.Employee) (err error) {
	errString := ""
	isCreated := false
	c.RLock()
	switch {
	case c == nil:
		errString += fmt.Sprintf("cache ReplaceId error: CacheEmployee can't be nil\n")
	case c.Cache == nil:
		errString += fmt.Sprintf("cache ReplaceId error: cache can't be nil\n")
	case employee == nil:
		errString += fmt.Sprintf("cache ReplaceId error: impossible to replace employee with nil value\n")
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
	if errString != "" {
		err = fmt.Errorf(errString)
	}
	return
}

func (c *CacheEmployee) UpsertId(employee *professions.Employee) (err error) {
	errString := ""
	isCreated := false
	c.RLock()
	switch {
	case c == nil:
		errString += fmt.Sprintf("cache UpsertId error: CacheEmployee can't be nil\n")
	case c.Cache == nil:
		errString += fmt.Sprintf("cache UpsertId error: cache can't be nil\n")
	case employee == nil:
		errString += fmt.Sprintf("cache UpsertId error: impossible to upsert employee with nil value\n")
	default:
		c.RUnlock()
		if _, isCreated, err = c.FindId(employee.ID); err != nil {
			return
		} else {
			if isCreated {
				err = c.ReplaceId(employee)
			} else {
				err = c.Add(*employee)
			}
		}
	}
	if errString != "" {
		fmt.Errorf(errString)
	}
	return
}

func (c *CacheEmployee) DeleteId(id uint64) (err error) {
	isCreated := false
	if _, isCreated, err = c.FindId(id); err != nil {
		return
	} else {
		if isCreated {
			delete(c.Cache, id)
		} else {
			err = fmt.Errorf("cache DeleteId error: impossible to delete element with id = %v because it doesn't exist\n", id)
		}
	}
	return
}

func (c *CacheEmployee) DeleteAll() (err error) {
	errString := ""
	switch {
	case c == nil:
		errString += fmt.Sprintf("cache DeleteAll error: CacheEmployee can't be nil\n")
	case c.Cache == nil:
		errString += fmt.Sprintf("cache DeleteAll error: cache can't be nil\n")
	default:
		c.Cache = make(map[uint64]professions.Employee)
	}
	return
}
