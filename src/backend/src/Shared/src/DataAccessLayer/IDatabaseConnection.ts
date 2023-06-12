export interface IDatabaseConnection {
	query: <T>(
		query: string[],
		params?: Array<Array<any>>,
		opts?: any,
	) => Promise<T>;
}
