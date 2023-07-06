package main

import (
	"fmt"
	"server/src/core/usecase"
	"server/src/infra/jsonrpc"
)

func main() {
	var accountRepository = jsonrpc.NewJsonRpc("http://localhost:8545")
	// var accountRepository = inmemory.InMemoryRepository{}

	var getAllAccounts = usecase.NewGetAllAccounts(accountRepository)
	fmt.Println(getAllAccounts.Exec())

	var getOneAccount = usecase.NewGetOneAccount(accountRepository)
	fmt.Println(getOneAccount.Exec("abc"))
}
