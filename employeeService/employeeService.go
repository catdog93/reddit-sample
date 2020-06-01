package employeeService

import (
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
	rep "github.com/catdog93/GoIT/repository"
	"gopkg.in/mgo.v2"
)

const (
	url            = "mongodb://127.0.0.1:2717"
	dbName         = "employees"
	collectionName = "employeesCollection"
	connectError   = "error occurs during establishing db connection: "
)

var results []rep.Obj = []rep.Obj{}
var result *rep.Obj

func CreateEmployee(employee *prof.Employee) error { // id front ????? //
	emp, err := ReadEmployeeIdd(employee.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if emp == nil {
		session, err := mgo.Dial(url)
		if err != nil {
			fmt.Println(connectError, err)
			return err
		}
		defer session.Close()

		err = rep.Create(session.DB(dbName).C(collectionName), employee)
	}
	return err
}

func ReadEmployeeIdd(id uint64) (*rep.Obj, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(connectError, err)
		return nil, err
	}
	defer session.Close()

	err = rep.ReadIdd(rep.ID{ID: id}, result, session.DB(dbName).C(collectionName))
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func ReadEmployeeId(id uint64) ([]rep.Obj, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(connectError, err)
		return nil, err
	}
	defer session.Close()

	err = rep.ReadId(rep.ID{ID: id}, results, session.DB(dbName).C(collectionName))
	if err != nil {
		fmt.Println(err)
	}
	return results, err
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
