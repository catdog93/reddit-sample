package professions

import (
	"strconv"
	"time"
)

const customDateFormat string = "02-Jan-2006"

type EmployeeService interface {
	GetEmployeeInfo() string
	GetEmployeePosition() string
}

type Employee struct {
	*Person    `json:"person"`
	HiringDate time.Time `json:"hiringDate,createdAt"`
	Salary     int       `json:"salary,omitempty"`
	FullTime   bool      `json:"fullTime,omitempty"`
	Position   string    `json:"position"`
	Role       Role      `json:"role"`
}

type Role string

const (
	EmployeeRole Role = "employee"
	BossRole     Role = "boss"
)

func (employee *Employee) GetEmployeeInfo() string {
	return employee.Person.GetPersonInfo() + "\nRole: " + string(employee.Role) + "\n" // + "\nHiring date: " + string(employee.HiringDate.Format(customDateFormat)) + ", salary: $ " + strconv.Itoa(employee.Salary) + "\nFull time: " + strconv.FormatBool(employee.FullTime) + "\n"
}

type EmployeeToPersonConverter interface {
	ConvertEmployeeToPerson() *Person
}

func (employee *Employee) ConvertEmployeeToPerson() *Person {
	return employee.Person
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

func (employee *Employee) GetEmployeePosition() string {
	return employee.GetPersonInfo() + ", position: " + employee.Position
}
