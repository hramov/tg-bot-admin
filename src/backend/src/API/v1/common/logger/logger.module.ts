import { Module } from '@nestjs/common';
import { CustomLoggerService } from './custom-logger.service';
import { AsyncLocalStorageModule } from '../asyncLocalStorage/asyncLocalStorage.module';
import { LOGGER } from '../constants';

@Module({
	imports: [AsyncLocalStorageModule],
	providers: [
		{
			provide: LOGGER,
			useClass: CustomLoggerService,
		},
	],
	exports: [LOGGER],
})
export class LoggerModule {}
