import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {OrderService} from "./order.service";
import {OrderController} from "./order.controller";
import {TypeOrmModule} from "@nestjs/typeorm";
import {OrderEntity} from "../../common/persistent/entity/order.entity";

@Module({
    imports: [LoggerModule, TypeOrmModule.forFeature([OrderEntity])],
    providers: [OrderService],
    controllers: [OrderController]
})
export class OrderModule {}
