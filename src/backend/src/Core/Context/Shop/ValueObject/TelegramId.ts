import {ValueObject} from "../../../../Shared/src/ValueObject/ValueObject";

export class TelegramName extends ValueObject {
    private readonly id: string;

    protected *getEqualityComponents() {
        yield this.id;
    }

}