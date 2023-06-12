import { Uuid } from 'src/Shared/src/ValueObject/Objects/Uuid';

describe('ValueObject', () => {
	describe('Check equality', () => {
		it('Not equal', () => {
			expect(new Uuid().equals(new Uuid())).toBe(false);
		});

		it('Equal', () => {
			const uuid = new Uuid();
			expect(uuid.equals(uuid)).toBe(true);
		});
	});
});
