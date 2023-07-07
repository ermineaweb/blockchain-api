package main

import (
	"server/src/core/adapter"
	"server/src/core/repository"
	"server/src/core/service"
	"server/src/infra/gin"
	"server/src/infra/jsonrpc"
)

func main() {
	var accountRepository repository.Account
	accountRepository = jsonrpc.NewJsonRpc("http://localhost:8545")
	// accountRepository = inmemory.InMemoryRepository{}

	var accountService = service.NewAccountService(accountRepository)

	var server adapter.Server
	// server := mux.NewMuxServer("0.0.0.0:5000", *accountService)
	server = gin.NewGinServer("0.0.0.0:5000", *accountService)
	server.Serve()
}
