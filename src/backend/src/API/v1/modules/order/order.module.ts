import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {OrderService} from "./order.service";
import {OrderController} from "./order.controller";

@Module({
    imports: [LoggerModule],
    providers: [OrderService],
    controllers: [OrderController]
})
export class OrderModule {}
