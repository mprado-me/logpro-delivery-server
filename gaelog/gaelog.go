package gaelog

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func PrintRequest(r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "%s", r)
}

func PrintString(r *http.Request, s string) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "%s", s)
}
