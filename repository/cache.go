package repository

import (
	"fmt"
	"github.com/catdog93/GoIT/professions"
	//"log"

	//"gopkg.in/mgo.v2"
	"sync"
)

type CacheEmployee struct {
	Cache map[uint64]professions.Employee
	sync.RWMutex
}

var Cache = CacheEmployee{
	Cache: make(map[uint64]professions.Employee),
}

func (c CacheEmployee) Add(docs ...professions.Employee) (err error) {
	errString := ""
	c.Lock()
	for _, v := range docs {
		if _, ok := c.Cache[v.ID]; ok {
			errString += fmt.Sprintf("Cache Add err : existed ID : %d, \n", v.ID)
		} else {
			c.Cache[v.ID] = v
		}
	}
	c.Unlock()
	if errString != "" {
		return fmt.Errorf(errString)
	}
	return
}

func (c CacheEmployee) FindId(id uint64) (employee professions.Employee, isCreated bool) {
	c.RLock()
	defer c.RUnlock()
	if _, ok := c.Cache[id]; ok {
		employee = c.Cache[id]
		isCreated = true
	}
	return
}

/*
func (c CacheEmployee) Update(employee professions.Employee) (err error) {
	c.Lock()
	c.Unlock()
}*/

func (c CacheEmployee) ReplaceId(employee professions.Employee) (err error) {
	id := employee.ID
	c.RLock()
	if _, isCreated := c.FindId(id); isCreated {
		c.Cache[id] = employee
	} else {
		err = fmt.Errorf("Cache ReplaceId err : not existed ID : %d, \n", id)
	}
	c.RUnlock()
	return
}

/*func (c CacheEmployee) UpsertId(id uint64, employee professions.Employee) (err error) {
	c.Lock()
	if _, isCreated := c.FindId(id); isCreated {
		c.Unlock()
		// Update
	} else {
		c.Unlock()
		err = c.Add(employee)
	}

}*/

/*
func (*cache) Update(selector interface{}, update interface{}, collection *mgo.Collection) (err error) {
	if collection != nil {
		err = collection.Update(selector, update)
	}
	return
}

func (*cache) Delete(selector interface{}, collection *mgo.Collection) (err error) {
	if collection != nil {
		err = collection.Remove(selector)
	}
	return
}

func (*cache) ReadId(ID uint64, collection *mgo.Collection) (resultQuery *mgo.Query) {
	if collection != nil && ID > 0 {
		resultQuery = collection.FindId(ID)
	}
	return
}

func (*cache) ReplaceId(ID uint64, update interface{}, collection *mgo.Collection) (err error) {
	if collection != nil && ID > 0 {
		err = collection.ReplaceId(ID, update)
	}
	return
}

func (*cache) DeleteId(ID uint64, collection *mgo.Collection) (err error) {
	if collection != nil && ID > 0 {
		err = collection.Remove(ID)
	}
	return
}

func (*cache) DeleteAll(selector interface{}, collection *mgo.Collection) (info *mgo.ChangeInfo, err error) {
	if collection != nil {
		info, err = collection.RemoveAll(selector)
	}
	return
}

func (cache) UpdateAll(selector interface{}, update interface{}, collection *mgo.Collection) (info *mgo.ChangeInfo, err error) {
	if collection != nil {
		info, err = collection.UpdateAll(selector, update)
	}
	return
}
*/
