package mux

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/src/core/service"
)

type getAccount struct {
	service service.Account
}

func (s getAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("getAccount")
	s.getAccount(rw, r)
}

func (p getAccount) getAccount(rw http.ResponseWriter, r *http.Request) {
	json := json.NewEncoder(rw)
	account, _ := p.service.Get("abc")
	json.Encode(account)
}
