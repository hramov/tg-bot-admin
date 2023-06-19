import { Module } from '@nestjs/common';
import {LoggerModule} from "./common/logger/logger.module";
import {ShopModule} from "./modules/shop/shop.module";
import {OrderModule} from "./modules/order/order.module";
import {ProductModule} from "./modules/product/product.module";
import {AuthModule} from "./modules/auth/auth.module";
import {PostgresModule} from "./common/persistent/postgres/postgres.module";
import {UserModule} from "./modules/user/user.module";

@Module({
  imports: [PostgresModule, LoggerModule, ShopModule, OrderModule, ProductModule, AuthModule, UserModule],
})
export class AppModule {}
