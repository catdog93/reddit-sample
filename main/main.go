package main

import (
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
)

/*
Создание Кеша  Работников (разных проффесий),  ключ - идентификатор ID.
Range по мапе, с целью опросить всех работников о их имени и должности.
*/

type ID int

func homeTask3() {
	employees := map[ID]prof.EmployeeService{ // cash's values has interface type
		0: prof.Astronaut{
			Employee: &prof.Employee{
				Person: &prof.Person{
					Name:     "Ivan",
					LastName: "Torch",
					Age:      25,
				},
				Position: "Captain",
			},
		},
		1: prof.Actor{
			Employee: &prof.Employee{
				Person: &prof.Person{
					Name:     "Will",
					LastName: "Smith",
					Age:      40,
				},
				Position: "Superhero",
			},
		},
		2: prof.SoftwareDeveloper{
			Employee: &prof.Employee{
				Person: &prof.Person{
					Name:     "Tim",
					LastName: "Dev",
					Age:      30,
				},
				Position: "Jun dev",
			},
		},
		4: &prof.Employee{
			Person: &prof.Person{
				Name:     "Abstract",
				LastName: "Employee",
				Age:      23,
			},
			Position: "None",
		},
	}

	fmt.Println()
	for key := range employees {
		fmt.Print(employees[key].GetEmployeePosition())
		fmt.Println()
	}
}

func main() {
	homeTask3()
}
