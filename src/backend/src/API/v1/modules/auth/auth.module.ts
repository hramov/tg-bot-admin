import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {AuthService} from "./auth.service";
import {AuthController} from "./auth.controller";

@Module({
    imports: [LoggerModule],
    providers: [AuthService],
    controllers: [AuthController]
})
export class AuthModule {}
