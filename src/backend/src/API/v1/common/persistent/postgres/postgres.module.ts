import { Module } from '@nestjs/common';
import {
	IPostgresConnOptions,
	PostgresStorage,
} from '../../../../../Infrastructure/Persistent/Storage/Postgres/Postgres.storage';
import {POSTGRES_STORAGE} from "../../constants";

@Module({
	providers: [
		{
			provide: POSTGRES_STORAGE,
			useFactory: () => {
				const options: IPostgresConnOptions = {
					host: process.env.PG_HOST,
					port: Number(process.env.PG_PORT),
					user: process.env.PG_USERNAME,
					password: process.env.PG_PASSWORD,
					database: process.env.PG_DATABASE,
				};
				return new PostgresStorage(options);
			},
		},
	],
	exports: [POSTGRES_STORAGE],
})
export class PostgresModule {}
