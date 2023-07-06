package jsonrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"server/src/core/entity"
	"time"
)

type JsonRpc struct {
	url    string
	client http.Client
}

type JsonRpcRequest struct {
	Id      int      `json:"id"`
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

type JsonRpcResponse struct {
	Id      int    `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result  []any  `json:"result"`
}

var accounts = []entity.Account{{Address: "abc", Balance: 15}, {Address: "def", Balance: 20}}

func NewJsonRpc(url string) JsonRpc {
	client := http.Client{Timeout: 5 * time.Second}
	return JsonRpc{url, client}
}

func (r JsonRpc) GetOne(address string) (entity.Account, error) {
	for _, account := range accounts {
		if account.Address == address {
			return account, nil
		}
	}
	return entity.Account{}, errors.New("account not found")
}

func (r JsonRpc) GetAll() []entity.Account {
	var request = JsonRpcRequest{Id: 1, JsonRpc: "2.0", Method: "eth_accounts", Params: make([]string, 0)}
	marshalled, _ := json.Marshal(request)

	req, _ := http.NewRequest("POST", r.url, bytes.NewReader(marshalled))
	req.Header.Set("Content-Type", "application/json")
	res, err := r.client.Do(req)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var response JsonRpcResponse
	_ = json.Unmarshal(body, &response)
	fmt.Printf("%+v\n", response)

	return accounts
}
