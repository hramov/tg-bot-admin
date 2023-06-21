import { Injectable } from '@nestjs/common';
import {UserDto} from "./dto/user.dto";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ProductSearchFilter} from "../../common/filters/product/search.filter";
import {InjectRepository} from "@nestjs/typeorm";
import {UserEntity} from "../../common/persistent/entity/user.entity";
import {Repository} from "typeorm";
import {CannotFindRoleError} from "../auth/error/CannotFindRole.error";
import {genSalt, hash} from "bcrypt";
import {RoleEntity} from "../../common/persistent/entity/role.entity";
import {RegisterDto} from "../auth/dto/register.dto";

@Injectable()
export class UserService {

    constructor(
        @InjectRepository(UserEntity) private readonly userRepository: Repository<UserEntity>,
        @InjectRepository(RoleEntity) private readonly roleRepository: Repository<RoleEntity>
    ) {}

    get(filters: ProductSearchFilter) {
        return Promise.resolve(undefined);
    }

    async getById(userId: Uuid) {
        return this.userRepository.findOne({
            where: {
                id: userId.toString(),
            },
            relations: {
                role: true,
                shop: true,
            },
        });
    }

    async getByTgName(tgName: string) {
        return this.userRepository.findOne({
            where: {
                tg_name: tgName,
            },
            relations: {
                role: true
            }
        });
    }

    async create(dto: RegisterDto) {
        const user = this.userRepository.create();
        user.tg_name = dto.tg_name;

        const role = await this.roleRepository.findOneBy({code: dto.role_code});
        if (!role || !role.id) {
            return new CannotFindRoleError();
        }

        user.role = role;

        const salt = await genSalt(5)
        user.password = await hash(dto.password, salt);

        await this.userRepository.save(user);
        return user;
    }

    delete(userId: Uuid) {
        return Promise.resolve(undefined);
    }

}
