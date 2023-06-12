import { Injectable } from '@nestjs/common';
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ShopDto} from "./dto/shop.dto";
import {ShopSearchFilter} from "../../common/filters/shop/search.filter";
import {IShopRepository} from "../../../../Core/Context/Shop/IShopRepository";

@Injectable()
export class ShopService {

    constructor(private readonly shopRepository: IShopRepository) {}


    get(filters: ShopSearchFilter) {
        return this.shopRepository.getByFilters(filters);
    }

    getById(shopId: Uuid) {
        return this.shopRepository.getById(shopId);
    }

    create(dto: ShopDto) {
        return Promise.resolve(undefined);
    }

    update(dto: ShopDto, shopId: Uuid) {
        return Promise.resolve(undefined);
    }

    delete(shopId: Uuid) {
        return Promise.resolve(undefined);
    }
}
