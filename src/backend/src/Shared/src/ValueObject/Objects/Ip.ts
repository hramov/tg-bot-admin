import {ValueObject} from "../ValueObject";

export class Ip extends ValueObject {
    public value: string;

    constructor(ip: string) {
        super();
        this.value = ip;
    }
    protected *getEqualityComponents(): IterableIterator<Object> {
        return undefined;
    }

    toString() {
        return String(this.value)
    }

}