import { Injectable } from '@nestjs/common';
import {ProductDto} from "./dto/product.dto";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ProductSearchFilter} from "../../common/filters/product/search.filter";

@Injectable()
export class ProductService {
    delete(productId: Uuid) {
        return Promise.resolve(undefined);
    }

    update(productId: Uuid) {
        return Promise.resolve(undefined);
    }

    create(dto: ProductDto) {
        return Promise.resolve(undefined);
    }

    getById(productId: Uuid) {
        return Promise.resolve(undefined);
    }

    get(filters: ProductSearchFilter) {
        return Promise.resolve(undefined);
    }
}
