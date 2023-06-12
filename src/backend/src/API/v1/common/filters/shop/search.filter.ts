import {Filter} from "../filters";
import {ShopDto} from "../../../modules/shop/dto/shop.dto";

export class ShopSearchFilter extends Filter {
    private readonly filter: ShopDto;

    constructor(query: string) {
        super(query);

        const filterMap = super.parse();
        this.filter.title = filterMap.get('title');
        this.filter.description = filterMap.get('description');
    }

    get(): ShopDto {
        return this.filter;
    }
}