import {Filter} from "../filters";
import {CreateShopDto} from "../../../modules/shop/dto/create-shop.dto";

export class ShopSearchFilter extends Filter {
    private readonly filter: CreateShopDto;

    constructor(query: string) {
        super(query);
    }

    get(): CreateShopDto {
        return this.filter;
    }
}