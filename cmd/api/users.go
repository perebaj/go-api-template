package main

import (
	"fmt"
	"net/http"
)

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new user")
}

func (app *application) showUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "show the details of a specific user")
}
