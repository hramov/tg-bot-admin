import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {RepositoryModule} from "../../common/persistent/repository/repository.module";
import {UserService} from "./user.service";
import {UserController} from "./user.controller";

@Module({
    imports: [LoggerModule, RepositoryModule],
    providers: [UserService],
    controllers: [UserController]
})
export class UserModule {}
