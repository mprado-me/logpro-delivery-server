package main

import (
	"admin"
	"github.com/gorilla/mux"
	"mob"
	"net/http"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/mobile/auth-user", mob.AuthUserHandler)
	r.HandleFunc("/website/admin/put-motoboy", admin.PutMotoboy)
	http.Handle("/", r)
}
