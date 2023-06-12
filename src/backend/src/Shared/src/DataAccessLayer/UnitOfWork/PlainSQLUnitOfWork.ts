import { IDatabaseConnection } from '../IDatabaseConnection';
import { QueryOptions } from '../QueryOptions';
import { IUnitOfWork } from './IUnitOfWork';

export class PlainSQLUnitOfWork implements IUnitOfWork {
	constructor(private readonly connection: IDatabaseConnection) {
		if (!connection) throw new Error('Connection is undefined');
	}

	async query<T>(
		query: string[],
		params?: Array<Array<string>>,
		opts?: QueryOptions,
	) {
		let result: T | Error;

		try {
			const result = await this.connection.query(query, params);
		} catch (_err: unknown) {
			result = _err as Error;
		}

		return result;
	}
}
