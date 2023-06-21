import { Module } from '@nestjs/common';
import {LoggerModule} from "./common/logger/logger.module";
import {AuthModule} from "./modules/auth/auth.module";
import {PostgresModule} from "./common/persistent/postgres/postgres.module";
import {UserModule} from "./modules/user/user.module";
import {ProductModule} from "./modules/product/product.module";
import {OrderModule} from "./modules/order/order.module";
import {ShopModule} from "./modules/shop/shop.module";

@Module({
  imports: [
    PostgresModule, LoggerModule, AuthModule, UserModule, ProductModule, OrderModule, ShopModule
  ],
})
export class AppModule {}
