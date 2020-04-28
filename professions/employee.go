package professions

import (
	"strconv"
	"time"
)

const customDateFormat string = "02-Jan-2006"

type EmployeeService interface {
	GetEmployeeInfo() string
}

type Employee struct {
	*Person    `json:"person"`
	HiringDate time.Time `json:"hiringDate,createdAt"`
	Salary     int       `json:"salary,omitempty"`
	FullTime   bool      `json:"fullTime,omitempty"`
}

func (employee Employee) GetEmployeeInfo() string {
	return employee.Person.GetPersonInfo() + "\nHiring date: " + string(employee.HiringDate.Format(customDateFormat)) + ", salary: $ " + strconv.Itoa(employee.Salary) + "\nFull time: " + strconv.FormatBool(employee.FullTime)
}

type PersonService interface {
	GetPersonInfo() string
}

type Person struct {
	Name        string `json:"name"`
	LastName    string `json:"lastName"`
	Age         int    `json:"age,omitempty"`
	Nationality string `json:"-"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
}

func (person Person) GetPersonInfo() string {
	return person.Name + " " + person.LastName + ", " + strconv.Itoa(person.Age) + " years old"
}
