import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {ProductService} from "./product.service";
import {ProductController} from "./product.controller";
import {TypeOrmModule} from "@nestjs/typeorm";
import {ShopEntity} from "../../common/persistent/entity/shop.entity";
import {ProductEntity} from "../../common/persistent/entity/product.entity";
import {CategoryEntity} from "../../common/persistent/entity/category.entity";

@Module({
    imports: [LoggerModule, TypeOrmModule.forFeature([ProductEntity, ShopEntity, CategoryEntity])],
    providers: [ProductService],
    controllers: [ProductController]
})
export class ProductModule {}
