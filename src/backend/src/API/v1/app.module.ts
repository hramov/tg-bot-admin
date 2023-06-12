import { Module } from '@nestjs/common';
import {LoggerModule} from "./common/logger/logger.module";
import {ShopModule} from "./modules/shop/shop.module";
import {OrderModule} from "./modules/order/order.module";
import {ProductModule} from "./modules/product/product.module";
import {AuthModule} from "./modules/auth/auth.module";

@Module({
  imports: [LoggerModule, ShopModule, OrderModule, ProductModule, AuthModule],
})
export class AppModule {}
