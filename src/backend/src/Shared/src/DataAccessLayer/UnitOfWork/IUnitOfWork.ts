import { QueryOptions } from '../QueryOptions';

export interface IUnitOfWork {
	query: <T>(
		query: string[],
		params?: Array<Array<any>>,
		opts?: QueryOptions,
	) => Promise<T | Error>;
}
