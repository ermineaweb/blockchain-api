class JsonRpcRepository implements Repository<Account> {
  private host: string;

  constructor(host: string) {
    this.host = host;
  }

  async get(address: string): Promise<Account> {
    console.log("repository get");

    const rawRes = await fetch(this.host, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ id: 1, jsonrpc: "2.0", method: "eth_accounts", params: [] }),
    });
    const jsonRes = await rawRes.json();
    console.log(jsonRes);

    return jsonRes;
  }

  async getAll(): Promise<Account[]> {
    console.log("repository getAll");

    const rawRes = await fetch(this.host, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ id: 1, jsonrpc: "2.0", method: "eth_accounts", params: [] }),
    });
    const jsonRes = await rawRes.json();

    return jsonRes.result;
  }
}

export default JsonRpcRepository;

/*
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
*/
