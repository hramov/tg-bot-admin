import {BaseEntity} from "../../../../Shared/src/BaseEntity";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {TelegramName} from "../ValueObject/TelegramId";
import {Address} from "../ValueObject/Address";
import {CoreProduct} from './Product';

export enum OrderStatus {
    Requested = 1,
    Purchased,
    Collected,
    Completed
}

export class Order extends BaseEntity<Uuid> {
    private readonly shop_id: Uuid;
    private readonly manager_id: Uuid;
    private readonly customer_tg_name: TelegramName;
    private readonly courier_name: string;

    private readonly delivery_address: Address;
    private readonly order_tracking_link: string;

    private readonly products: CoreProduct[];
    private readonly status: OrderStatus;

    private readonly order_time: Date;
    private readonly expected_time: Date;
    private readonly sending_time: Date;
    private readonly delivery_time: Date;
}