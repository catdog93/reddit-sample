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
	cache := make(Obj)
	cache["person1"] = &prof.Person{
		Name:     "Tim",
		LastName: "Ku",
	}
	cache["person2"] = &prof.Person{
		Name:     "Mat",
		LastName: "Tom",
	}
	cache["person3"] = &prof.Person{
		Name:     "Nick",
		LastName: "Cool",
	}

	session, err := mgo.Dial("mongodb://127.0.0.1:2717")
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		collection := session.DB("test1").C("testCollection")
		ai.Connect(collection)

		empService := rep.ProfessionsService{Collection: collection}

		if err := empService.Create(cache["person1"]); err != nil { // BSON field 'insert.documents.0' is the wrong type 'array', expected type 'object'
			log.Fatal(err)
		} else {
			fmt.Println(empService.Read(Obj{}))
		}
	}
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
