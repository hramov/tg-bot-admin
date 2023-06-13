import { Module } from '@nestjs/common';
import {PostgresModule} from "../postgres/postgres.module";
import {LoggerModule} from "../../logger/logger.module";
import {SHOP_REPOSITORY} from "./repository.constants";
import {ILogger} from "../../../../../Core/ICore";
import {IDatabaseConnection} from "../../../../../Infrastructure/Persistent/Storage/IDatabaseConnection";
import {IPostgresQueryOptions} from "../../../../../Infrastructure/Persistent/Storage/Postgres/IPostgresQueryOptions";
import {ShopRepository} from "./shop.repository";
import {LOGGER, POSTGRES_STORAGE} from "../../constants";

@Module({
	imports: [PostgresModule, LoggerModule],
	providers: [
        {
            provide: SHOP_REPOSITORY,
			useFactory: (logger: ILogger, storage: IDatabaseConnection<IPostgresQueryOptions>) => {
				return new ShopRepository(logger, storage)
			},
			inject: [LOGGER, POSTGRES_STORAGE],
        }
	],
	exports: [SHOP_REPOSITORY],
})
export class RepositoryModule {}
