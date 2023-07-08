import express from "express";
import service from "../core/service/account";
import AccountService from "../core/service/account";

class ExpressServer implements Server {
  private service: AccountService;
  private app;

  constructor(service: service) {
    this.service = service;
    this.app = express();

    this.app.use((req, res, next) => {
      console.log(`request: ${req.method} ${req.path}`);
      return next();
    });

    this.app.get("/api/account/:id", async (req, res) => {
      const id = req.params.id;
      const account = await this.service.get(id);
      return res.status(200).json(account);
    });

    this.app.get("/api/accounts", async (req, res) => {
      const accounts = await this.service.getAll();
      return res.status(200).json(accounts);
    });
  }

  serve(host: string, port: number): void {
    this.app.listen(port, host, () => {
      console.log("server started");
    });
  }
}

export default ExpressServer;
