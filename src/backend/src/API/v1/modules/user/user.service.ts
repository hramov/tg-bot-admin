import { Injectable } from '@nestjs/common';
import {UserDto} from "./dto/user.dto";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ProductSearchFilter} from "../../common/filters/product/search.filter";

@Injectable()
export class UserService {
    delete(productId: Uuid) {
        return Promise.resolve(undefined);
    }

    update(productId: Uuid) {
        return Promise.resolve(undefined);
    }

    create(dto: UserDto) {
        return Promise.resolve(undefined);
    }

    getById(productId: Uuid) {
        return Promise.resolve(undefined);
    }

    get(filters: ProductSearchFilter) {
        return Promise.resolve(undefined);
    }
}
