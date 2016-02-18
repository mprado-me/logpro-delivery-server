package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"mob"
	"net/http"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/mobile/auth-user", mob.AuthUserHandler)
	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World! 2")
}
