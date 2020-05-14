package main

import (
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
)

type Obj map[string]interface{}

var result []interface{}

func main() {
	/*var cash Obj = make(Obj)
	cash["student1"] = prof.Person{Name: "Tim", LastName: "123"}
	cash["student2"] = prof.Person{Name: "Mat", LastName: "Tom"}
	cash["student3"] = prof.Person{Name: "Nick", LastName: "Cool"}
	cash["student4"] = prof.Person{Name: "Rob", LastName: "Pop"}
	*/

	session, err := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()
	if err != nil {
		fmt.Println("error ", err) // OR panic(error)
	} else {
		ai.Connect(session.DB("test1").C("testCollection"))
		slice := []prof.Person{
			prof.Person{ID: ai.Next("testCollection"), Name: "AAAAAA", LastName: "123"},
			prof.Person{ID: ai.Next("testCollection"), Name: "BBBBBB", LastName: "Tom"},
		}
		var sliceDB []interface{}
		for index := range slice {
			sliceDB = append(sliceDB, slice[index])
		}
		err := session.DB("test1").C("testCollection").Insert(sliceDB...)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
		/*	err = session.DB("test1").C("testCollection").Find(Obj{"_id":1}).All(&result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
			err = session.DB("test1").C("testCollection").RemoveId(1)
			if err != nil {
				fmt.Println(err)
			}*/
		/*err = session.DB("test1").C("testCollection").Update(Obj{"_id":2}, Obj{"$set":Obj{"name":"OLEG"}}) // ! ! !
		if err != nil {
			fmt.Println(err)
		}*/

		err = session.DB("test1").C("testCollection").Find(Obj{}).All(&result)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}

}

/*func sum(array []int, flag bool) (sum int) {
	if flag { // sum even numbers
		for _, value := range array {
			if value%2 == 0 {
				sum += value
			}
		}
	} else { // sum odd index number
		for _, value := range array {
			if value%2 != 0 {
				sum += value
			}
		}
	}
	fmt.Println(sum)
	return
}

func sumWG(array []int, flag bool) (sum int) {
	if flag { // sum even numbers
		for _, value := range array {
			if value%2 == 0 {
				sum += value
			}
		}
	} else { // sum odd index number
		for _, value := range array {
			if value%2 != 0 {
				sum += value
			}
		}
	}
	fmt.Println(sum)
	wg.Done()
	return
}

var wg = sync.WaitGroup{}

func main() {
	array := []int{0, 1, 2, 44, 55}

	go sum(array, true)
	go sum(array, false)
	wg.Add(2)
	go sumWG(array, false)
	wg.Wait()
	wg.Add(1)
	go sumWG(array, true)
	wg.Wait()
}

/*counter++
fmt.Println("main ", counter)
if !flag {
	flag = true
	f := Funcer{
		func1: main,
		func2: main,
	}
	fmt.Println(f.ReturnMethods())
}*/
/*
type funcer interface {
	ReturnMethods() []*func()
}

type Funcer struct {
	func1 func()
	func2 func()
}

//ReturnMethods
func (f *Funcer) ReturnMethods() (funcSlice []*func()) {
	if f != nil {
		funcSlice = append(funcSlice, &f.func1)
		funcSlice = append(funcSlice, &f.func2)
	}
	return
}*/
