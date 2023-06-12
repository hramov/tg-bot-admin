import { ILoggerWriter } from './Logger';

export class ConsoleLogger implements ILoggerWriter {
	async write(msg: string, stack?: any): Promise<boolean | Error> {
		console.log(msg);
		if (stack) console.log(stack);
		return true;
	}
}
