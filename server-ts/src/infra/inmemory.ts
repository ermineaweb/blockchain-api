class InMemoryRepository implements Repository<Account> {
  static accounts: Account[] = [
    { address: "abc", balance: 22 },
    { address: "def", balance: 15 },
  ];

  get(address: string): Promise<Account> {
    const account = InMemoryRepository.accounts.find((elem) => elem.address === address);
    return new Promise((resolve) => resolve(account));
  }

  getAll(): Promise<Account[]> {
    const accounts = InMemoryRepository.accounts;
    return new Promise((resolve) => resolve(accounts));
  }
}

export default InMemoryRepository;
