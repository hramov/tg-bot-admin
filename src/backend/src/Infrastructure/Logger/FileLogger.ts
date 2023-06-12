import { join } from 'path';
import { FileHelper } from '../FileHelper/FileHelper';
import { ILoggerWriter } from './Logger';

export class FileLogger implements ILoggerWriter {
	private readonly fileHelper = new FileHelper();
	private readonly logsPath = process.env.LOGS_PATH;

	private createLogsFileName() {
		const dt = new Date();
		return dt.getFullYear() + '-' + dt.getMonth() + '-' + dt.getDate();
	}

	async write(msg: string, stack?: any): Promise<boolean | Error> {
		const result = await this.fileHelper.append(
			join(this.logsPath, this.createLogsFileName()),
			msg + ',\n',
		);
		if (result instanceof Error) {
			return result;
		}
		return true;
	}
}
