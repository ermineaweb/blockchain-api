package mux

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/src/core/service"
)

type getAllAccounts struct {
	service service.Account
}

func (s getAllAccounts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("getAllAccounts")
	s.getAccounts(rw, r)
}

func (p getAllAccounts) getAccounts(rw http.ResponseWriter, r *http.Request) {
	json := json.NewEncoder(rw)
	json.Encode(p.service.GetAll())
}
