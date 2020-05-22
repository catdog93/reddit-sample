package professionsService

import (
	"encoding/json"
	"fmt"
	prof "github.com/catdog93/GoIT/professions"
	rep "github.com/catdog93/GoIT/repository"
	"io/ioutil"
	"net/http"
)

type PersonService struct {
}

type PersonCRUD string

var (
	CreatePerson PersonCRUD = "/persons/create/"
	ReadPerson   PersonCRUD = "/persons/read/"
	UpdatePerson PersonCRUD = "/persons/update/"
	DeletePerson PersonCRUD = "/persons/delete/"
)

func (p *PersonService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(rw, "", err)
			fmt.Println(err)
		}
		tempPerson := &prof.Person{}
		profService := &rep.ProfessionsService{}

		personCRUD := PersonCRUD(r.URL.Path)
		switch personCRUD {
		case CreatePerson:
			if err := json.Unmarshal(body, tempPerson); err != nil {
				fmt.Fprintf(rw, "", err)
				fmt.Println(err)
			} else {
				if e, err := tempPerson.ConvertPersonToEmployee(); err != nil {
					fmt.Fprintf(rw, "", err)
					fmt.Println(err)
				} else {
					if err := profService.Create(*e); err != nil {
						fmt.Fprintf(rw, "", err)
						fmt.Println(err)
					} else {
						result := profService.ReadId(tempPerson.ID)
						fmt.Fprintf(rw, "", result)
						fmt.Println(result)
					}
				}
			}
		}
	/*case ReadPerson:

	case UpdatePerson:

	case DeletePerson:*/

	default:
		fmt.Fprintf(rw, "", r.URL.Path)
		fmt.Println(r.URL.Path)
	}
	fmt.Fprintf(rw, "", r.PostForm)
}
