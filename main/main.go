package main

import (
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
	rep "github.com/catdog93/GoIT/repository"
	/*ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"*/
	"log"
)

func main() {
	e1 := prof.Employee{
		ID: 1,
		Person: prof.Person{
			ID:       1,
			Name:     "Nick",
			LastName: "Cool",
		},
		Salary: 100,
	}
	e2 := prof.Employee{
		ID: 101,
		Person: prof.Person{
			ID:       1,
			Name:     "Vasya",
			LastName: "Cat",
		},
		Salary: 100,
	}
	if err := rep.Cache.Add(e1); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(rep.Cache.Cache)
	}
	emp, isCreated := rep.Cache.FindId(1)
	fmt.Println(emp, isCreated)

	rep.Cache.ReplaceId(e2)

	emp, isCreated = rep.Cache.FindId(1)
	fmt.Println(emp, isCreated)
	/*fmt.Println()
	rep.Cache.ReplaceId(1, e2)

	emp, isCreated = rep.Cache.FindId(1)
	fmt.Println(emp, isCreated)*/

	/*isCreated = rep.Cache.FindId(2)
	fmt.Printf("document has already created: %v\n", isCreated)*/

	/*
		session, err := mgo.Dial("mongodb://127.0.0.1:2717")
		defer session.Close()
		if err != nil {
			log.Fatal(err)
		} else {
			collection := session.DB("test1").C("testCollection")
			ai.Connect(collection) // bson.NewObjectId()

			p := prof.Person{
				ID: ai.Next(collection.Name),
				Name:     "Nick",
				LastName: "Cool",
			}

			empService := rep.ProfessionsService{Collection: collection}

			if info, err := empService.DeleteAll(rep.Obj{}); err != nil {
				log.Fatal(err, info)
			}
			if err := empService.Add(p); err != nil {
				log.Fatal(err)
			} else {
				if err := empService.Read(rep.Obj{}).All(&result); err != nil {

				} else {
					fmt.Println(result)
				}
			}
		}*/
}

/*sliceCache := []interface{}{
	prof.Employee{
		Person: &prof.Person{
			Name:     "Nick",
			LastName: "Cool",
		},
	},
	prof.Employee{
		Person: &prof.Person{
			Name:     "Mat",
			LastName: "Tom",
		},
	},
	prof.Employee{
		Person: &prof.Person{
			Name:     "Tim",
			LastName: "Ku",
		},
	},
}*/
