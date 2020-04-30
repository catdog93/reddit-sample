package main

import (
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
	"time"
)

/*Создать пару указателей на существующие структуры, вызвать панику.
Сменить в структурах вложение структур на вложение указателей.
Написать методы для указателя на человека. (GetPersonInfo, GetContacts, SetAge, NewPersonBorned)
Вызвать эти методы для работников.*/

func homeTask2() {
	var employeeIvan *prof.Employee
	var astronaut *prof.Astronaut

	//panic("pointers are nill")

	employeeIvan.GetPersonInfo() // panic nil pointer or invalid memory address
	employeeIvan.GetEmployeeInfo()

	astronaut.GetAstronautInfo()
}

func main() {

}

func test1() {
	person := &prof.Person{ // creating Person instance
		Name:        "Ivan",
		LastName:    "Pavlov",
		Age:         20,
		Nationality: "Golang",
	}
	employee := &prof.Employee{ // creating Employee instance
		Person: person,
		//HiringDate: time.Date(2019, 11, 1)
		HiringDate: time.Now(),
		Salary:     3000,
		FullTime:   true,
	}

	fmt.Println()
	fmt.Println(employee.GetEmployeeInfo())
	fmt.Println(employee.Person)

	specialSkills := []prof.AustronautSpecialSkill{prof.SystemEngineer, prof.Captain} // creating slices for special fields of Astronaut instance
	spacecrafts := []prof.SpacecraftModel{prof.ChinaShuttle, prof.SovietShuttle}
	astronaut := &prof.Astronaut{ // creating Astronaut instance
		Employee:      employee,
		SpecialSkills: specialSkills,
		Spacecrafts:   spacecrafts,
		Experience:    5,
	}

	fmt.Println(astronaut.GetAstronautInfo())
}

/*
	a := 3
	b := 7

	fmt.Println()
	fmt.Println(sum(&a, &b))
	fmt.Println(a, b)
	fmt.Println()
	fmt.Println(sum2(a, b))
	fmt.Println(a, b)

func sum(a, b *int) int {
	*a = 2
	*b = 6
	return *a + *b
}

func sum2(a, b int) int {
	a = 2
	b = 3
	return a + b
}
*/
