package main

import (
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
	"sync"
)

/*
1. Создать функцию которая приводит работников к человеку(значение типов)

2. Создать кеши начальников и работников не приступать к опросу работников пока не будут перебраны все начальники.
Перебирать кеши в рутинах, конкурентно. С RWMutex.
*/

var wg sync.WaitGroup
var mux sync.RWMutex

type ID uint32

func PrintEmployeesCash(cash map[ID]*prof.Employee) {
	mux.Lock()
	for key := range cash {
		fmt.Println(cash[key].GetEmployeeInfo())
	}
	// Decrement the counter when the goroutine completes.
	defer wg.Done()
	mux.Unlock()
}

func main() {
	// 2. Асинхронно перебираем используя WaitGroup & RWMutex
	chiefs := map[ID]*prof.Employee{
		1: &prof.Employee{
			Person: &prof.Person{
				Name:     "Ivan",
				LastName: "Torch",
				Age:      25,
			},
			Position: "NASA's Head",
			Role:     "boss",
		},

		2: &prof.Employee{
			Person: &prof.Person{
				Name:     "Will",
				LastName: "Smith",
				Age:      40,
			},
			Position: "Producer",
			Role:     "boss",
		},

		3: &prof.Employee{
			Person: &prof.Person{
				Name:     "Tim",
				LastName: "Dev",
				Age:      30,
			},
			Position: "Manager",
			Role:     "boss",
		},
	}
	employees := map[ID]*prof.Employee{
		1: &prof.Employee{
			Person: &prof.Person{
				Name:     "Boris",
				LastName: "cool",
				Age:      40,
			},
			Position: "Superhero",
			Role:     "employee",
		},

		2: &prof.Employee{
			Person: &prof.Person{
				Name:     "Tim",
				LastName: "Dev",
				Age:      30,
			},
			Position: "Jun dev",
			Role:     "employee",
		},

		3: &prof.Employee{
			Person: &prof.Person{
				Name:     "Abstract",
				LastName: "Employee",
				Age:      23,
			},
			Position: "None",
			Role:     "employee",
		},
	}

	wg.Add(1)
	go PrintEmployeesCash(chiefs)
	wg.Wait()
	wg.Add(1)
	go PrintEmployeesCash(employees)
	wg.Wait()

	/*
		// 1. Создать функцию которая приводит работников к человеку(значение типов)
			e := &Employee{
				Person: &Person{
					Name:     "Abstract",
					LastName: "Employee",
					Age:      23,
				},
				Position: "None",
			}
			fmt.Println(e.ConvertEmployeeToPerson())*/
}
