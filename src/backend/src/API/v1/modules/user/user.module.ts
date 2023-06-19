import { Module } from '@nestjs/common';
import {LoggerModule} from "../../common/logger/logger.module";
import {UserService} from "./user.service";
import {UserController} from "./user.controller";
import {TypeOrmModule} from "@nestjs/typeorm";
import {UserEntity} from "../../common/persistent/entity/user/user.entity";
import {RoleEntity} from "../../common/persistent/entity/user/role.entity";

@Module({
    imports: [LoggerModule, TypeOrmModule.forFeature([UserEntity, RoleEntity])],
    providers: [UserService],
    controllers: [UserController]
})
export class UserModule {}
