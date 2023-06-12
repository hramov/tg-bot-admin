import { ConsoleLogger } from './ConsoleLogger';
import {ILogger} from "../../Core/ICore";

export interface ILoggerWriter {
	write(msg: string, stack?: any): Promise<boolean | Error>;
}

export enum LogLevel {
	DEBUG = 'debug',
	INFO = 'info',
	WARN = 'warn',
	ERROR = 'error',
}

export class Logger implements ILogger {
	private readonly consoleLogger = new ConsoleLogger();

	constructor(private readonly context: string) {}

	public debug(msg: string, context?: string, opts?: any) {
		const message = this.createLogMessage('debug', msg, context, opts);
		const outputs = this.outputFactory(LogLevel.DEBUG);

		for (const output of outputs) {
			output.write(message).catch((err) => console.error(err));
		}
	}

	public log(msg: string, context?: string, opts?: any) {
		const message = this.createLogMessage('info', msg, context, opts);
		const outputs = this.outputFactory(LogLevel.INFO);

		for (const output of outputs) {
			output.write(message).catch((err) => console.error(err));
		}
	}

	public warn(msg: string, context?: string, opts?: any) {
		const message = this.createLogMessage('warn', msg, context, opts);
		const outputs = this.outputFactory(LogLevel.WARN);

		for (const output of outputs) {
			output.write(message).catch((err) => console.error(err));
		}
	}

	public error(msg: string, context?: string, stack?: any, opts?: any) {
		const message = this.createLogMessage('error', msg, context, opts);
		const outputs = this.outputFactory(LogLevel.ERROR);

		for (const output of outputs) {
			output.write(message).catch((err) => console.error(err));
		}
	}

	private outputFactory(level: LogLevel): ILoggerWriter[] {
		// TODO return all output from config
		return [this.consoleLogger];
	}

	private createLogMessage(
		type: string,
		message: string,
		context: string,
		opts?: any,
	) {
		return JSON.stringify({
			ts: new Date().toISOString(),
			type: type,
			context: context || this.context,
			message: message,
			...opts,
		});
	}
}
