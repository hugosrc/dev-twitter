package routes

import (
	"fmt"
	"net/http"
)

// CreateUser is the http route to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
