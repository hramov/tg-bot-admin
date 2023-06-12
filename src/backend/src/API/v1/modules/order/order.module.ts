import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {RepositoryModule} from "../../common/persistent/repository/repository.module";
import {OrderService} from "./order.service";
import {OrderController} from "./order.controller";

@Module({
    imports: [LoggerModule, RepositoryModule],
    providers: [OrderService],
    controllers: [OrderController]
})
export class OrderModule {}
