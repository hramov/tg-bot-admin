import { createHash } from 'crypto';

export abstract class ValueObject {
	protected abstract getEqualityComponents(): IterableIterator<Object>;

	public equals<T extends ValueObject>(obj: T): boolean {
		const thisEqualityGenerator = this.getEqualityComponents();
		const objEqualityGenerator = obj.getEqualityComponents();

		while (true) {
			const thisValue = thisEqualityGenerator.next();
			const objValue = objEqualityGenerator.next();

			if (thisValue.done) return true;

			if (thisValue.value !== objValue.value) return false;
		}
	}

	public getHashCode(): string {
		return String(createHash('sha256'));
	}
}
