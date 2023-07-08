interface Repository<T> {
  get(address: string): Promise<T>;
  getAll(): Promise<T[]>;
}
