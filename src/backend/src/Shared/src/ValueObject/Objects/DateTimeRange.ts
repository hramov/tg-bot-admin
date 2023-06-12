import * as moment from 'moment';
import { ValueObject } from '../ValueObject';
import { DateTimeDuration } from './DateTimeDuration';

export class DateTimeRange extends ValueObject {
	public start: Date;
	public end: Date;
	public duration: DateTimeDuration;

	constructor(start: Date, end?: Date, duration?: string) {
		super();
		this.start = start;
		if (end) this.end = end;

		if (!end && duration) {
			this.duration = new DateTimeDuration(duration);
			this.end = this.endFromStartAndDuration(this.start, this.duration);
		}
	}

	protected override *getEqualityComponents(): IterableIterator<Object> {
		yield this.start;
		yield this.end;
	}

	public durationInMinutes(): number {
		return moment
			.duration(moment(this.end).diff(moment(this.start)))
			.asMinutes();
	}

	private endFromStartAndDuration(
		start: Date,
		duration: DateTimeDuration,
	): Date {
		return null;
	}

	public newDuration(newDuration: string) {
		return new DateTimeRange(this.start, null, newDuration);
	}

	public newEnd(newEnd: Date) {
		return new DateTimeRange(this.start, newEnd);
	}

	public newStart(newStart: Date) {
		return new DateTimeRange(newStart, this.end);
	}

	public static createOneDayRange(day: Date) {
		const newDay = new Date(day);

		return new DateTimeRange(
			day,
			new Date(newDay.setDate(day.getDate() + 1)),
		);
	}

	public static createOneWeekRange(day: Date) {
		const newDay = new Date(day);

		return new DateTimeRange(
			day,
			new Date(newDay.setDate(day.getDate() + 7)),
		);
	}

	public overlaps(dateTimeRange: DateTimeRange) {
		return this.start < dateTimeRange.end && this.end > dateTimeRange.start;
	}
}
