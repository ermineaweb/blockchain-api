## Documentations

https://ethereum.github.io/execution-apis/api-documentation/

https://open-rpc.org/getting-started

## JsonRpc

request example

```bash
curl http://localhost:8545 -H "Content-Type: application/json" \
--data '{"id":1,"jsonrpc":"2.0","method":"eth_accounts","params":[]}'
```

response example

```json
{
  "id": 1,
  "jsonrpc": "2.0",
  "result": ["0xd8e4fa854b1619bdcd014b98528db369c492e661", "0x94897a59dff2d75b38033e5609d2a1f6101cf88a"]
}
```
