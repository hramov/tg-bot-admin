import {ValueObject} from "../../../../Shared/src/ValueObject/ValueObject";

export const enum Currency {
    RUB = 'rub',
    USB = 'usd'
}

export class Price extends ValueObject {

    private readonly amount: number;
    private readonly currency: Currency;

    protected *getEqualityComponents() {
        yield this.amount;
        yield this.currency;
    }

}