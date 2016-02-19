package admin

import (
	"fmt"
	"mtb"
	"net/http"
)

func PutMotoboy(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	email := values.Get("email")
	password := values.Get("password")
	motoboy := mtb.NewMotoboy(email, password)
	err := mtb.Put(r, motoboy)
	if err != nil {
		fmt.Fprintf(w, "FAIL")
	} else {
		fmt.Fprintf(w, "SUCCESS")
	}
}
