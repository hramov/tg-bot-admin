import {ValueObject} from "../../../../Shared/src/ValueObject/ValueObject";

export class Address extends ValueObject {
    private readonly address: string;

    protected *getEqualityComponents() {
        yield this.address;
    }

}