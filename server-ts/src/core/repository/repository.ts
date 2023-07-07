interface Repository<T> {
  get(id: string): T;
  getAll(): T[];
}
