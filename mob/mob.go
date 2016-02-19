package mob

import (
	"fmt"
	"gaelog"
	"mtb"
	"net/http"
)

func AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	email := values.Get("email")
	password := values.Get("password")
	gaelog.PrintString(r, "email = "+email)
	gaelog.PrintString(r, "password = "+password)
	motoboy := &mtb.Motoboy{}
	err := mtb.Get(r, email, motoboy)
	if err != nil {
		fmt.Fprintf(w, "EMAIL_NOT_REGISTERED")
		return
	}
	if password != motoboy.Password {
		fmt.Fprintf(w, "INCORRECT_PASSWORD")
		return
	}
	fmt.Fprintf(w, "SUCCESS")
}
