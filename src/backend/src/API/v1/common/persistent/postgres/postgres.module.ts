import { Module } from '@nestjs/common';
import {TypeOrmModule} from "@nestjs/typeorm";

@Module({
	imports: [
		TypeOrmModule.forRootAsync({
			useFactory: () => ({
				type: 'postgres',
				host: process.env.PG_HOST,
				port: Number(process.env.PG_PORT),
				username: process.env.PG_USERNAME,
				password: process.env.PG_PASSWORD,
				database: process.env.PG_DATABASE,
				entities: [__dirname + '../entity/**/*.entity.ts'],
				autoLoadEntities: true,
				synchronize: true,
			})
		})
	],
})
export class PostgresModule {}
