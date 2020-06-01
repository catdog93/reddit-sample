package controller

import (
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
	service "github.com/catdog93/GoIT/service"
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
	DeleteEmployee     EmployeeCRUD = "/delete/id"
	DeleteAllEmployees EmployeeCRUD = "/delete/all"
)

type ID struct {
	ID uint64 `json:"id" form:"id"`
}

func ReadAll(context *gin.Context) {
	obj, err := service.ReadAllEmployees()
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusNotFound, err.Error())
		return
	}
	context.JSON(http.StatusOK, obj)
}

func ReadAllPost(context *gin.Context) {
	obj, err := service.ReadAllEmployees()
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusNotFound, err.Error())
		return
	}
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
	obj, err := service.ReadEmployeeId(id.ID)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusNotFound, err.Error())
		return
	}
	context.JSON(http.StatusOK, obj)
}

func ReadIdPost(context *gin.Context) {
	id := ID{}
	err := context.BindJSON(&id)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	obj, err := service.ReadEmployeeId(id.ID)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusNotFound, err.Error())
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
	id, err := service.CreateEmployee(&employee)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusCreated, id)
}

func ReplaceId(context *gin.Context) {
	employee := prof.Employee{Person: &prof.Person{}}
	err := context.BindJSON(&employee)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	err = service.UpdateEmployeeId(&employee)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusOK, err.Error())
		return
	}
}

func DeleteId(context *gin.Context) {
	id := ID{}
	err := context.BindJSON(&id)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	err = service.DeleteEmployeeId(id.ID)
	if err != nil {
		fmt.Println(err)
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
}
