package mux

import (
	"fmt"
	"log"
	"net/http"
	"server/src/core/service"
)

type muxServer struct {
	host    string
	service service.Account
}

func NewMuxServer(host string, service service.Account) *muxServer {
	return &muxServer{host, service}
}

func (s *muxServer) Serve() {
	server := http.NewServeMux()
	server.Handle("/api/accounts", getAllAccounts{s.service})
	server.Handle("/api/account", getAccount{s.service})

	fmt.Println("start http server")
	if err := http.ListenAndServe(s.host, server); err != nil {
		log.Fatal("failed to start server")
	}
}
