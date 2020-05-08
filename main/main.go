package main

import (
	"fmt"
	"strconv"
	//prof "github.com/catdog93/GoIT/professions"
)

/*
 Создать мапу типа [ID] interface{}
Написать Функцию которая будет принимать кеш и возвращать типы значений каждого элемента Кэша.
Поигратся с го рутинами. Реализовать алгоритм Луна
*/
/*
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
*/
func main() {
	cardNumber := 5375414118690212
	getReverseInt(cardNumber)
}

const digitsNumber = 16

func getReverseInt(number int) int {
	var stringsSlice []string
	var string string
	digit := 0
	for number > 0 {
		digit = number % 10
		string = strconv.Itoa(digit)
		stringsSlice = append(stringsSlice, string)
		number = number / 10
	}
	/*if result, error := strconv.Atoi(strings.Join(stringsSlice,"")); error == nil {
		return result
	} else {
		panic(error)
	}*/
	fmt.Println(stringsSlice)
	//fmt.Println(strconv.Atoi(strings.Join(stringsSlice,"")))
	return 0
}

func moonAlgorithmCheckCardNumber(cardNumber int) bool {
	getReverseInt(cardNumber)
	return true
}

func conversion(i interface{}) {
	switch value := i.(type) {
	case int:
		if value, ok := i.(int); ok {
			fmt.Println(value)
		}
	case string:
		if value, ok := i.(string); ok {
			fmt.Println(value)
		}
	default:
		fmt.Printf("I don't know about type %T!\n", value)
	}
}
