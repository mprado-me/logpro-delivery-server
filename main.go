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
	r.HandleFunc("/mobile/chat-msg-write", mob.WriteChatMsgHandler)
	r.HandleFunc("/mobile/chat-msg-read", mob.ReadChatMsgHandler)
	r.HandleFunc("/website/admin/put-motoboy", admin.PutMotoboy)
	http.Handle("/", r)
}
