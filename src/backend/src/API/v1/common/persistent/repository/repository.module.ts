import { Module } from '@nestjs/common';
import {PostgresModule} from "../postgres/postgres.module";
import {LoggerModule} from "../../logger/logger.module";

@Module({
	imports: [PostgresModule, LoggerModule],
	providers: [
	],
	exports: [],
})
export class RepositoryModule {}
