package professionsService

import (
	"encoding/json"
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
	rep "github.com/catdog93/GoIT/repository"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
)

type PersonService struct {
}

type EmployeeCRUD string

var (
	CreateEmployee  EmployeeCRUD = "/employees/create/"
	ReadEmployee    EmployeeCRUD = "/employees/read/id"
	ReplaceEmployee EmployeeCRUD = "/employees/replace/id"
	DeleteEmployee  EmployeeCRUD = "/employees/delete/id"

	ReadEmployees   EmployeeCRUD = "/employees/read/all"
	DeleteEmployees EmployeeCRUD = "/employees/delete/all"
)

type ID struct {
	ID uint64 `json:"id" bson:"_id"`
}

type ReplaceIdRequestBody struct {
	*prof.Employee `json:"employee" bson:"employee"`
}

var results []rep.Obj

func (p *PersonService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Fprintf(rw, "", err)
			fmt.Println("error ServeHTTP ", err)
		}
		errString := "error: "
		id := &ID{}

		session, err := mgo.Dial("mongodb://127.0.0.1:27017")
		defer session.Close()
		if err != nil {
			fmt.Fprintf(rw, "", err)
			fmt.Println("error ServeHTTP ", err)
		} else {
			collection := session.DB("test1").C("testCollection")
			ai.Connect(collection)
			empService := rep.ProfessionsService{Collection: collection}

			employeeCRUD := EmployeeCRUD(r.URL.Path)

			switch employeeCRUD {
			case CreateEmployee:
				inputEmployee := &prof.Employee{}
				if err := json.Unmarshal(body, inputEmployee); err != nil {
					fmt.Fprintf(rw, "", err, err)
					fmt.Println(errString, CreateEmployee, err)
				} else {
					inputEmployee.ID = ai.Next(empService.Collection.Name)
					inputEmployee.Person.ID = ai.Next(empService.Collection.Name)
					if err := empService.Create(*inputEmployee); err != nil {
						fmt.Fprintf(rw, "", err)
						fmt.Println(errString, CreateEmployee, err)
					} else {
						rw.WriteHeader(http.StatusCreated)
					}
				}
			case ReadEmployees:
				if err := empService.ReadAll(&results); err != nil {
					fmt.Fprintf(rw, "", err)
					fmt.Println(errString, ReadEmployees, err)
				} else {
					if results == nil {
						rw.WriteHeader(http.StatusNotFound)
					} else {
						if bytes, err := json.Marshal(results); err != nil {
							fmt.Fprintf(rw, "", err)
							fmt.Println(errString, ReadEmployees, err)
						} else {
							fmt.Fprintf(rw, "%s", bytes)
							fmt.Println(bytes)
						}
					}
				}
			case ReadEmployee:
				if err := json.Unmarshal(body, &id); err != nil {
					fmt.Fprintf(rw, "", err)
					fmt.Println(errString, ReadEmployee, err)
				} else {
					if err := empService.ReadId(id.ID, &results); err != nil {
						fmt.Fprintf(rw, "", err)
						fmt.Println(errString, ReadEmployee, err)
					} else {
						if results == nil {
							rw.WriteHeader(http.StatusNotFound)
						} else {
							if bytes, err := json.Marshal(results); err != nil {
								fmt.Fprintf(rw, "", err)
								fmt.Println(errString, ReadEmployee, err)
							} else {
								fmt.Fprintf(rw, "%s", bytes)
								fmt.Println(bytes)
							}
						}
					}
				}
			case ReplaceEmployee:
				r := ReplaceIdRequestBody{}
				if err := json.Unmarshal(body, &r); err != nil {
					fmt.Fprintf(rw, "", err)
					fmt.Println(errString, ReplaceEmployee, err)
				} else {
					if err := empService.UpdateId(r.ID, r.Employee); err != nil {
						rw.WriteHeader(http.StatusNotFound)
						fmt.Fprintf(rw, "", err)
						fmt.Println(errString, ReplaceEmployee, err)
					}
				}
			case DeleteEmployee:
				if err := json.Unmarshal(body, &id); err != nil {
					fmt.Fprintf(rw, "", err)
					fmt.Println(errString, DeleteEmployee, err)
				} else {
					if err := empService.DeleteId(id.ID); err != nil { // ? action
						fmt.Fprintf(rw, "", err)
						fmt.Println(errString, DeleteEmployee, err)
					}
				}
			case DeleteEmployees:
				if _, err := empService.DeleteAll(); err != nil {
					fmt.Fprintf(rw, "", err)
					fmt.Println(errString, DeleteEmployees, err)
				}
			default:
				fmt.Fprintf(rw, "", r.URL.Path)
				fmt.Println(r.URL.Path)
			}
		}
	}
}
