package mob

import (
	"encoding/json" //Temp
	"fmt"
	"kind"
	"mem/short"
	"msg"
	"mtb"
	"net/http"
	"proc/chat"
	"user"
)

// TODO: Usar InAuthMsg, OutAuthMsg e procauth.Process()
// TODO: Não passar o login por query, pois apesar do https encriptar as informacoes de login,
// toda url requisitada é salva no log do servidor
func AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	email := values.Get("email")
	password := values.Get("password")

	motoboy := &mtb.Motoboy{}
	err := mtb.Get(r, email, motoboy)
	if err != nil {
		fmt.Fprintf(w, "{\"auth_ok\": false, \"reason\": \"email\"}")
		return
	}
	if password != motoboy.Password {
		fmt.Fprintf(w, "{\"auth_ok\": false, \"reason\": \"password\"}")
		return
	}
	fmt.Fprintf(w, "{\"auth_ok\": true}")
}

// TODO: Autenticar
func WriteChatMsgHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	email := values.Get("email")
	password := values.Get("password")
	text := values.Get("text")

	login := user.NewLogin(email, password, kind.MOTOBOY)

	inChatMsg := msg.InChat{
		In: msg.In{
			Sender: login,
		},
		Text: text,
	}
	prochat.Process(inChatMsg)
}

// TODO: Autenticar
func ReadChatMsgHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	email := values.Get("email")
	outMsgs, ok := smem.GetMsgsTo(email)
	if ok {
		jsonMsg, err := json.Marshal(outMsgs)
		if err != nil {
			panic("Response message json encode fail")
		}
		smem.ClearMsgsTo(email)
		w.Write(jsonMsg)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
