import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {ProductService} from "./product.service";
import {ProductController} from "./product.controller";

@Module({
    imports: [LoggerModule],
    providers: [ProductService],
    controllers: [ProductController]
})
export class ProductModule {}
