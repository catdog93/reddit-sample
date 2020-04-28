package main

import (
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
	"time"
)

func main() {
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
	specialSkills := []prof.AustronautSpecialSkill{prof.SystemEngineer, prof.Captain} // creating slices for special fields of Astronaut instance
	spacecrafts := []prof.SpacecraftModel{prof.ChinaShuttle, prof.SovietShuttle}
	astronaut := &prof.Astronaut{ // creating Astronaut instance
		Employee:      employee,
		SpecialSkills: specialSkills,
		Spacecrafts:   spacecrafts,
		Experience:    5,
	}

	fmt.Println(astronaut.GetAstronautInfo())
	//fmt.Println(employee.GetEmployeeInfo())
	//fmt.Println(person.GetPersonInfo())
}
