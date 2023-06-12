import * as pgPromise from 'pg-promise';
import {IClient} from 'pg-promise/typescript/pg-subset';
import {DatabaseError} from "../../../../Core/Error/Database.error";
import QueryResultError = pgPromise.errors.QueryResultError;
import { IDatabaseConnection } from '../IDatabaseConnection';
import { IPostgresQueryOptions } from './IPostgresQueryOptions';

export interface IPostgresConnOptions {
	host: string;
	port: number;
	database: string;
	user: string;
	password: string;
}

export class PostgresStorage implements IDatabaseConnection<IPostgresQueryOptions> {
	private readonly conn: pgPromise.IDatabase<IClient>;

	constructor(connOptions: IPostgresConnOptions) {
		this.conn = pgPromise()(connOptions);
	}

	public async query<T>(
		sql: string,
		values?: any[],
		options?: IPostgresQueryOptions,
	): Promise<T[] | DatabaseError> {
		try {
			const data = await this.conn.many<T>(sql, values);
			return data;
		} catch (_err: unknown) {
			const err = _err as QueryResultError;
			return new DatabaseError(err.message + ' ' + PostgresStorage.cleanErrorMessage(err.query));
		}
	}

	public async queryTx<T>(
		sql: string[],
		values?: any[][],
		options?: IPostgresQueryOptions,
	): Promise<T | DatabaseError> {

		if (values) {
			if (sql.length !== values.length) {
				return new DatabaseError('Количество запросов не соответствует количеству параметров')
			}
		}

		try {
			const data = await this.conn.tx<T>(async (t) => {
				let data = null;
				for (let i = 0; i < sql.length; i++) {
					data = await t.one<T>(sql[i], values[i])
				}
				return data;
			});
			return data;
		} catch (_err: unknown) {
			const err = _err as QueryResultError;
			return new DatabaseError(err.message + ' ' + PostgresStorage.cleanErrorMessage(err.query));
		}
	}

	public async queryOne<T>(
		sql: string,
		values?: any[],
		options?: IPostgresQueryOptions,
	): Promise<T | DatabaseError> {
		try {
			const data = await this.conn.one<T>(sql, values);
			return data;
		} catch (_err: unknown) {
			const err = _err as QueryResultError;
			return new DatabaseError(err.message + ' ' + PostgresStorage.cleanErrorMessage(err.query));
		}
	}

	private static cleanErrorMessage(message: string) {
		if (!message) return '';
		return message.replace(/ +(?= )/g,'').replace(/(\r\n|\n|\r)/gm, "").trim();
	}
}
