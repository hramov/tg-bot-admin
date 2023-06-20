import {Filter} from "../filters";
import {ShopDto} from "../../../modules/shop/dto/shop.dto";

export class ShopSearchFilter extends Filter {
    private readonly filter: ShopDto;

    constructor(query: string) {
        super(query);
    }

    get(): ShopDto {
        return this.filter;
    }
}