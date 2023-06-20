import { Injectable } from '@nestjs/common';
import {UserDto} from "./dto/user.dto";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ProductSearchFilter} from "../../common/filters/product/search.filter";
import {InjectRepository} from "@nestjs/typeorm";
import {UserEntity} from "../../common/persistent/entity/user.entity";
import {Repository} from "typeorm";

@Injectable()
export class UserService {

    constructor(@InjectRepository(UserEntity) private readonly userRepository: Repository<UserEntity>) {}

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

    save(dto: UserDto) {
        return Promise.resolve(undefined);
    }

    delete(userId: Uuid) {
        return Promise.resolve(undefined);
    }

}
