import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {RepositoryModule} from "../../common/persistent/repository/repository.module";
import {AuthService} from "./auth.service";
import {AuthController} from "./auth.controller";

@Module({
    imports: [LoggerModule, RepositoryModule],
    providers: [AuthService],
    controllers: [AuthController]
})
export class AuthModule {}
