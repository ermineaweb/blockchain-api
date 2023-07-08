package jsonrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func NewJsonRpc(url string) JsonRpc {
	client := http.Client{Timeout: 5 * time.Second}
	return JsonRpc{url, client}
}

func (r JsonRpc) Get(address string) (entity.Account, error) {
	// TODO
	return entity.Account{Address: "abc", Balance: 15}, errors.New("account not found")
}

func (r JsonRpc) GetAll() []entity.Account {
	jsonRequest := JsonRpcRequest{Id: 1, JsonRpc: "2.0", Method: "eth_accounts", Params: make([]string, 0)}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(jsonRequest)

	req, err := http.NewRequest("POST", r.url, body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err.Error())
		return []entity.Account{}
	}

	res, err := r.client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return []entity.Account{}
	}

	response := new(JsonRpcResponse)
	json.NewDecoder(res.Body).Decode(&response)

	var accounts []entity.Account
	for _, res := range response.Result {
		accounts = append(accounts, entity.Account{Address: res.(string), Balance: 15})
	}

	return accounts
}
