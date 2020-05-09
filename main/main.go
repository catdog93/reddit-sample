package main

import (
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
	"reflect"
	"strconv"
)

/*
 Создать мапу типа [ID] interface{}
Написать Функцию которая будет принимать кеш и возвращать типы значений каждого элемента Кэша.
Поигратся с го рутинами. Реализовать алгоритм Луна
*/

type ID int

func cashOfEmptyInterfaceType() {
	employees := map[ID]interface{}{ // cash's values has interface{} type
		0: &prof.Astronaut{
			Employee: &prof.Employee{
				Person: &prof.Person{
					Name:     "Ivan",
					LastName: "Torch",
					Age:      25,
				},
				Position: "Captain",
			},
		},
		1: &prof.Actor{
			Employee: &prof.Employee{
				Person: &prof.Person{
					Name:     "Will",
					LastName: "Smith",
					Age:      40,
				},
				Position: "Superhero",
			},
		},
		2: &prof.SoftwareDeveloper{
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
	fmt.Println(mapContainsTypes(employees))
}

func mapContainsTypes(cash map[ID]interface{}) (types []interface{}) {
	for key := range cash {
		types = append(types, reflect.TypeOf(cash[key]))
	}
	return
}

func main() {
	// reflect.TypeOf() for each element of map[ID]interface{}
	cashOfEmptyInterfaceType()

	// Moon algorithm
	cardNumbers := []cardNumber{5375414118690212, 378282246310005, 5019717010103742, 76009244561, 4222222222222, 2222990905257051}
	for _, value := range cardNumbers {
		fmt.Println(moonAlgorithmCheckCardNumber(value))
	}
}

type cardNumber int // alias type

func moonAlgorithmCheckCardNumber(cardNumber cardNumber) bool {
	if len(strconv.Itoa(int(cardNumber))) > 12 { // cardNumbers has length > 12
		var intSlice []int
		digit, sum := 0, 0
		for index := 0; cardNumber > 0; index++ { // % 10 return digit from the end
			digit = (int(cardNumber)) % 10
			intSlice = append(intSlice, digit)
			cardNumber = cardNumber / 10 // /10 makes integer shorter per 1 digit at the end
			if index%2 != 0 {            // even digits * 2
				intSlice[index] *= 2
				if intSlice[index] > 9 { // subtract 9 from any number > 9
					intSlice[index] -= 9
				}
			}
			sum += intSlice[index] // sum all gotten digits
		}
		return sum%10 == 0
	} else {
		return false
	}
}
