import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {ShopService} from "./shop.service";
import {ShopController} from "./shop.controller";
import {TypeOrmModule} from "@nestjs/typeorm";
import { ShopEntity } from "../../common/persistent/entity/shop.entity";
import {ProductEntity} from "../../common/persistent/entity/product.entity";
import {CategoryEntity} from "../../common/persistent/entity/category.entity";
import {UserEntity} from "../../common/persistent/entity/user.entity";
import {OrderEntity} from "../../common/persistent/entity/order.entity";

@Module({
    imports: [LoggerModule, TypeOrmModule.forFeature([ShopEntity, ProductEntity, OrderEntity, CategoryEntity, UserEntity])],
    providers: [ShopService],
    controllers: [ShopController]
})
export class ShopModule {}
