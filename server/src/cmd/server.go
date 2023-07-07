package main

import (
	"server/src/core/service"
	"server/src/infra/gin"
	"server/src/infra/jsonrpc"
)

func main() {
	// accountRepository = inmemory.InMemoryRepository{}
	accountRepository := jsonrpc.NewJsonRpc("http://localhost:8545")

	var accountService = service.NewAccountService(accountRepository)

	// server := mux.NewMuxServer("0.0.0.0:5000", *accountService)
	server := gin.NewGinServer("0.0.0.0:5000", *accountService)
	server.Serve()
}
