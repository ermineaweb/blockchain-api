class AccountService {
  private repository: Repository<Account>;

  constructor(repository: Repository<Account>) {
    this.repository = repository;
  }

  get(address: string): Promise<Account> {
    return this.repository.get(address);
  }

  getAll(): Promise<Account[]> {
    return this.repository.getAll();
  }
}

export default AccountService;
