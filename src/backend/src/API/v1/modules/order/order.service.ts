import { Injectable } from '@nestjs/common';
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {OrderDto} from "./dto/order.dto";
import {OrderSearchFilter} from "../../common/filters/order/search.filter";
import {InjectRepository} from "@nestjs/typeorm";
import {OrderEntity} from "../../common/persistent/entity/order.entity";
import {Repository} from "typeorm";

@Injectable()
export class OrderService {
    constructor(
        @InjectRepository(OrderEntity) private readonly orderRepository: Repository<OrderEntity>
    ) {}

    get(filters: OrderSearchFilter) {
        return this.orderRepository.find({ relations: { products: true }});
    }

    getById(orderId: Uuid) {
        return this.orderRepository.findOne({ where: {id: orderId.toString()}, relations: { products: true }});
    }

    create(dto: OrderDto) {
        return Promise.resolve(undefined);
    }

    cancel(orderId: Uuid) {
        return Promise.resolve(undefined);
    }
}
