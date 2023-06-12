export class InMemoryStorage {
	private state = new Map<string, any>();

	public set<T>(key: string, value: T): void {
		this.state.set(key, value);
	}

	public get(key: string) {
		return this.state.get(key);
	}
}
