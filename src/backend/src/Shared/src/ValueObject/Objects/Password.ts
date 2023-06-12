import {ValueObject} from "../ValueObject";

export class Password extends ValueObject{
    constructor(private readonly password: string) {
        super();
    }

    protected *getEqualityComponents(): IterableIterator<Object> {
        yield this.password
    }

    public toString() {
        return this.password;
    }
}