package main

import (
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
	rep "github.com/catdog93/GoIT/repository"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
	"log"
)

var (
	result []interface{}
)

type Obj map[string]interface{}

func main() {
	cash := make(Obj)
	cash["employee1"] = prof.Employee{
		Person: &prof.Person{
			Name:     "Tim",
			LastName: "Ku",
		},
	}
	cash["employee2"] = prof.Employee{
		Person: &prof.Person{
			Name:     "Mat",
			LastName: "Tom",
		},
	}
	cash["employee3"] = prof.Employee{
		Person: &prof.Person{
			Name:     "Nick",
			LastName: "Cool",
		},
	}

	session, err := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		collection := session.DB("test1").C("testCollection")
		ai.Connect(collection)

		fmt.Println(collection.Find(Obj{}))

		empService := rep.ProfessionsService{Collection: collection}

		if err := empService.Create(cash); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(empService.Read(Obj{}))
		}
	}
}
