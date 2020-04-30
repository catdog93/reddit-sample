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

func (employee *Employee) GetEmployeeInfo() string {
	employee.Name = "123"
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

func (person *Person) GetPersonInfo() string {
	return person.Name + " " + person.LastName + ", " + strconv.Itoa(person.Age) + " years old"
}

func (person *Person) GetContacts() string {
	return person.Phone + ", " + person.Email
}

func (person *Person) SetAge(age int) bool {
	if person != nil && age >= 0 {
		person.Age = age
		return true
	}
	return false
}

func (person *Person) NewPersonBorned(name, lastName, nationality string) bool {
	if person == nil {
		person = new(Person)
		person.Name = name
		person.LastName = lastName
		person.Nationality = nationality
		return true
	}
	return false
}
