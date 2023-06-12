import { ICronConfig, ICronConfigItem } from '../Core/ICore';

export const CronConfig: ICronConfig = new Map<string, ICronConfigItem>([
		['Job', {
				name: 'Job',
				enabled: true,
				cron: '*/10 * * * * *',
				params: {
					foo: 'bar',
				},
				runOnce: true,
		}]
	]
)