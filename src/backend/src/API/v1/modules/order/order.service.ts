import { Injectable } from '@nestjs/common';
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {OrderDto} from "./dto/order.dto";
import {OrderSearchFilter} from "../../common/filters/order/search.filter";

@Injectable()
export class OrderService {

    get(filters: OrderSearchFilter) {
        return Promise.resolve(undefined);
    }

    cancel(orderId: Uuid) {
        return Promise.resolve(undefined);
    }

    create(dto: OrderDto) {
        return Promise.resolve(undefined);
    }

    getById(orderId: Uuid) {
        return Promise.resolve(undefined);
    }
}
