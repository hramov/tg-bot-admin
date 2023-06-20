import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {AuthService} from "./auth.service";
import {AuthController} from "./auth.controller";
import {TypeOrmModule} from "@nestjs/typeorm";
import {UserEntity} from "../../common/persistent/entity/user.entity";
import {RoleEntity} from "../../common/persistent/entity/role.entity";
import {JwtModule} from "@nestjs/jwt";
import {APP_GUARD} from "@nestjs/core";
import {AuthGuard} from "./auth.guard";
import {RolesGuard} from "./roles.guard";
import {secret} from "./auth.constants";

@Module({
    imports: [LoggerModule,
        JwtModule.register({
            global: true,
            secret: secret, // TODO move to config
            signOptions: { expiresIn: '30m' },
        }),
        TypeOrmModule.forFeature([UserEntity, RoleEntity])],
    providers: [AuthService, {
        provide: APP_GUARD,
        useClass: AuthGuard,
    }, {
        provide: APP_GUARD,
        useClass: RolesGuard,
    },],
    controllers: [AuthController]
})
export class AuthModule {}
