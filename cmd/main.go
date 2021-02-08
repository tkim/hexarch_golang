// using CRUD
// routing is under http using rest
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com\tkim\hexarch_golang\tree\main\pkg\http\rest"
	"github.com\tkim\hexarch_golang\tree\main\pkg\storage"
)

func main() {

	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatalln("error while setting up storage:", err)
	}

	log.println(r.GetAllNames())

	fmt.Println("starting server on port 8080...")
	router := rest.InitHandlers()
	log.Fatal(http.ListenAndServe(":8080", router))
}