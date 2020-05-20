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
	slice := []prof.Employee{
		prof.Employee{
			ID: 1,
			Person: prof.Person{
				ID:       1,
				Name:     "Nick",
				LastName: "Cool",
			},
			Salary: 100,
		},
		prof.Employee{
			ID: 2,
			Person: prof.Person{
				ID:       1,
				Name:     "Vasya",
				LastName: "Cat",
			},
			Salary: 100,
		},
	}
	e2 := prof.Employee{
		ID: 1,
		Person: prof.Person{
			ID:       1,
			Name:     "OOP",
			LastName: "TRUE",
		},
		Salary: 100,
	}
	e3 := prof.Employee{
		ID: 1,
		Person: prof.Person{
			ID:       1,
			Name:     "Wit",
			LastName: "Weather",
		},
		Salary: 100,
	}
	e4 := prof.Employee{
		ID: 3,
		Person: prof.Person{
			ID:       1,
			Name:     "Upsert",
			LastName: "Do",
		},
		Salary: 10,
	}

	if err := rep.Cache.Add(slice...); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(rep.Cache.Cache)
	}
	if emp, isCreated, err := rep.Cache.FindId(e2.ID); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(emp, isCreated)
	}
	if err := rep.Cache.ReplaceId(&e2); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(rep.Cache.Cache)
	}
	if err := rep.Cache.UpsertId(&e3); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(rep.Cache.Cache)
	}
	if err := rep.Cache.UpsertId(&e4); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(rep.Cache.Cache)
	}
	/*if err := rep.Cache.DeleteId(4); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(rep.Cache.Cache)
	}*/
	if err := rep.Cache.DeleteAll(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(rep.Cache.Cache)
	}
}
