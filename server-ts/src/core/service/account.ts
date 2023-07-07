class AccountService {
  private repository: Repository<Account>;

  constructor(repository: Repository<Account>) {
    this.repository = repository;
  }
}
