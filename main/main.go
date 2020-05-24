package main

import (
	ps "github.com/catdog93/GoIT/professionsService"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", &ps.PersonService{})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
