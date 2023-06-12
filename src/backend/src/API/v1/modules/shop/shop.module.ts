import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {RepositoryModule} from "../../common/persistent/repository/repository.module";
import {ShopService} from "./shop.service";
import {ShopController} from "./shop.controller";

@Module({
    imports: [LoggerModule, RepositoryModule],
    providers: [ShopService],
    controllers: [ShopController]
})
export class ShopModule {}
