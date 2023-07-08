import AccountService from "../core/service/account";
import ExpressServer from "../infra/express";
import InMemoryRepository from "../infra/inmemory";
import JsonRpcRepository from "../infra/jsonrpc";

const accountRepository = new JsonRpcRepository("http://127.0.0.1:8545");
// const accountRepository = new InMemoryRepository();
const accountService = new AccountService(accountRepository);
const server = new ExpressServer(accountService);

server.serve("localhost", 4000);
