package employeeController

import (
	"fmt"
	service "github.com/catdog93/GoIT/employeeService"
	prof "github.com/catdog93/GoIT/professions"
	rep "github.com/catdog93/GoIT/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmployeeCRUD string

var (
	RelativePathEmployees EmployeeCRUD = "/employees"

	CreateEmployee     EmployeeCRUD = "/create/"
	ReadAllEmployees   EmployeeCRUD = "/read/all"
	ReadEmployee       EmployeeCRUD = "/read/id"
	ReplaceEmployee    EmployeeCRUD = "/replace/id"
	DeleteAllEmployees EmployeeCRUD = "/delete/all"
	DeleteEmployee     EmployeeCRUD = "/delete/id"
)

type ID struct {
	ID uint64 `json:"id" form:"id"`
}

type ReplaceIdRequestBody struct {
	*prof.Employee `json:"employee" bson:"employee"`
}

var results []rep.Obj

//func Create(context *gin.Context) {
//	inputEmployee := &prof.Employee{}
//	err := context.BindJSON(inputEmployee)
//	if err != nil {
//		fmt.Println(err)
//		context.String(http.StatusBadRequest, "", err)
//		return
//	}
//	// service
//	err = service.CreateEmployee(inputEmployee)
//	if err != nil {
//		fmt.Println(err)
//		//context.String(http.StatusInternalServerError, "", err)
//		context.String(http.StatusBadRequest, err.Error())
//		return
//	}
//	context.JSON(201, inputEmployee)
//}

func ReadIdPost(context *gin.Context) {
	id := ID{}
	err := context.BindJSON(&id)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	obj, err := service.ReadEmployeeId(id.ID)
	context.JSON(http.StatusOK, obj)
}

func ReadId(context *gin.Context) {
	id := ID{}
	err := context.BindQuery(&id)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	obj, err := service.ReadEmployeeIdd(id.ID)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusNotFound, err.Error()) // or 500
		return
	}
	context.JSON(http.StatusOK, obj)
}

func Create(context *gin.Context) {
	employee := prof.Employee{Person: &prof.Person{}}
	err := context.BindJSON(&employee)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	err = service.CreateEmployee(&employee)
	context.String(http.StatusCreated, "")
}

func ReadAll(c *gin.Context) {
	//if err := empService.ReadAll(&results); err != nil {
	//	fmt.Fprintf(rw, "", err)
	//	fmt.Println(errString, ReadAllEmployees, err)
	//} else {
	//	if results == nil {
	//		rw.WriteHeader(http.StatusNotFound)
	//	} else {
	//		if bytes, err := json.Marshal(results); err != nil {
	//			fmt.Fprintf(rw, "", err)
	//			fmt.Println(errString, ReadAllEmployees, err)
	//		} else {
	//			fmt.Fprintf(rw, "%s", bytes)
	//			fmt.Println(bytes)
	//		}
	//	}
	//}
}

/*if r.Method == "POST" {
body, err := ioutil.ReadAllEmployees(r.Body)
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
	empService := rep.employeeRepository{Collection: collection}

	employeeCRUD := EmployeeCRUD(r.URL.Path)*/

/*
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
			if err := empService.DeleteEmployee(id.ID); err != nil { // ? action
				fmt.Fprintf(rw, "", err)
				fmt.Println(errString, DeleteEmployee, err)
			}
		}
	case DeleteAllEmployees:
		if _, err := empService.DeleteAllEmployees(); err != nil {
			fmt.Fprintf(rw, "", err)
			fmt.Println(errString, DeleteAllEmployees, err)
		}
	default:
		fmt.Fprintf(rw, "", r.URL.Path)
		fmt.Println(r.URL.Path)
	}
}
}*/
