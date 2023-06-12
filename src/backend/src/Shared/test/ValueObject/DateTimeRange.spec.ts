import { DateTimeRange } from '../../ValueObject/Objects/DateTimeRange';

describe('DateTimeRange', () => {
	describe('createOneDayRange', () => {
		it('expected values', () => {
			const range = DateTimeRange.createOneDayRange(new Date());
			expect(range.start.getDate()).toBe(new Date().getDate());
			expect(range.end.getDate()).toBe(new Date().getDate() + 1);
		});
	});

	describe('createOneWeekRange', () => {
		it('expected values', () => {
			const range = DateTimeRange.createOneWeekRange(new Date());
			expect(range.start.getDate()).toBe(new Date().getDate());
			expect(range.end.getDate()).toBe(new Date().getDate() + 7);
		});
	});

	describe('durationInMinutes', () => {
		it('expected values', () => {
			const start = new Date();
			const dateTimeRange = DateTimeRange.createOneDayRange(start);
			expect(dateTimeRange.durationInMinutes()).toBe(24 * 60);
		});
	});

	describe('endFromStartAndDuration', () => {
		it('expected values', () => {
			const start = new Date();
			const dateTimeRange = DateTimeRange.createOneDayRange(start);
			expect(dateTimeRange.durationInMinutes()).toBe(24 * 60);
		});
	});
});
