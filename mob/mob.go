package mob

import (
	"fmt"
	"net/http"
)

func AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "1")
}
