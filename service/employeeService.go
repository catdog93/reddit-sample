package service

import (
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
	rep "github.com/catdog93/GoIT/repository"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
)

const (
	url            = "mongodb://127.0.0.1:2717"
	dbName         = "employees"
	collectionName = "employeesCollection"
	connectError   = "error occurs during establishing db connection: "
)

var results = []rep.Obj{}
var result = &rep.Obj{}

func ReadAllEmployees() ([]rep.Obj, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(connectError, err)
		return nil, err
	}
	defer session.Close()

	err = rep.ReadAll(results, session.DB(dbName).C(collectionName))
	if err != nil {
		fmt.Println(err)
	}
	return results, err
}

func ReadEmployeeId(id uint64) (*rep.Obj, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(connectError, err)
		return nil, err
	}
	defer session.Close()

	err = rep.ReadId(id, result, session.DB(dbName).C(collectionName))
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func CreateEmployee(employee *prof.Employee) (uint64, error) {
	//emp, err := ReadEmployeeId(employee.ID)
	//if err != nil {
	//	fmt.Println(err)
	//	return 0, err
	//}
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(connectError, err)
		return 0, err
	}
	defer session.Close()

	//if emp != nil {
	//	ai.Connect(session.DB(dbName).C(collectionName))
	//	employee.ID = ai.Next(collectionName)
	//}
	ai.Connect(session.DB(dbName).C(collectionName))
	err = rep.Create(session.DB(dbName).C(collectionName), employee)
	return employee.ID, err
}

func UpdateEmployeeId(employee *prof.Employee) error {
	emp, err := ReadEmployeeId(employee.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(connectError, err)
		return err
	}
	defer session.Close()

	if emp == nil {
		err = rep.Create(session.DB(dbName).C(collectionName), employee)
	} else {
		err = rep.UpdateId(employee.ID, employee, session.DB(dbName).C(collectionName))
	}
	return err
}

func DeleteEmployeeId(id uint64) error {
	emp, err := ReadEmployeeId(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if emp == nil {
		return nil
	}
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(connectError, err)
		return err
	}
	defer session.Close()

	err = rep.DeleteId(id, session.DB(dbName).C(collectionName))
	return err
}
