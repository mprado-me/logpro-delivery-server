// TODO: Armazenar os motoboys cadastrados em lmem
package mtb

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

const (
	kind = "mtb"
)

type RegisteredData struct {
	Email    string
	Password string
}

type Motoboy struct {
	RegisteredData
}

func NewMotoboy(email string, password string) *Motoboy {
	motoboy := new(Motoboy)
	motoboy.Email = email
	motoboy.Password = password
	return motoboy
}

func Put(r *http.Request, motoboy *Motoboy) error {
	ctx := appengine.NewContext(r)
	key := datastore.NewKey(ctx, kind, motoboy.Email, 0, nil)
	_, err := datastore.Put(ctx, key, motoboy)
	return err
}

func Get(r *http.Request, email string, motoboy *Motoboy) error {
	ctx := appengine.NewContext(r)
	key := datastore.NewKey(ctx, kind, email, 0, nil)
	return datastore.Get(ctx, key, motoboy)
}

func Delete(r *http.Request, email string) error {
	ctx := appengine.NewContext(r)
	key := datastore.NewKey(ctx, kind, email, 0, nil)
	return datastore.Delete(ctx, key)
}
