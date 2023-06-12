import {BaseEntity} from "../../../../Shared/src/BaseEntity";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {Product} from "./Product";
import {TelegramId} from "../ValueObject/TelegramId";
import {Address} from "../ValueObject/Address";

export enum OrderStatus {
    Requested = 1,
    Purchased,
    Collected,
    Completed
}

export class Order extends BaseEntity<Uuid> {
    private readonly shopId: Uuid;
    private readonly customerId: TelegramId;
    private readonly deliveryAddress: Address;

    private readonly products: Product[];
    private readonly status: OrderStatus;

    private readonly date_ordered: Date;
    private readonly date_delivered: Date;
}