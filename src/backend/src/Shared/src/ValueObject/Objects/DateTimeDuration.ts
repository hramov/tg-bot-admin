import { ValueObject } from '../ValueObject';

export const enum DateTimeDurationUnit {
	MILLISECOND = 'ms',
	SECOND = 's',
	MINUTE = 'm',
	HOUR = 'h',
	DAY = 'd',
	MONTH = 'mo',
	YEAR = 'y',
}

export class DateTimeDuration extends ValueObject {
	public duration: number;
	public unit: DateTimeDurationUnit;
	constructor(private readonly durationString: string) {
		super();
		this.parseDurationString();
	}

	private parseDurationString() {
		this.duration = 0;
		this.unit = DateTimeDurationUnit.DAY;
	}

	protected *getEqualityComponents(): IterableIterator<Object> {
		throw new Error('Method not implemented.');
	}
}
