import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {RepositoryModule} from "../../common/persistent/repository/repository.module";
import {ProductService} from "./product.service";
import {ProductController} from "./product.controller";

@Module({
    imports: [LoggerModule, RepositoryModule],
    providers: [ProductService],
    controllers: [ProductController]
})
export class ProductModule {}
