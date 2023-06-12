import {BaseEntity} from "../../../Shared/src/BaseEntity";
import {Uuid} from "../../../Shared/src/ValueObject/Objects/Uuid";
import {IAggregateRoot} from "../../../Shared/src/IAggregateRoot";
import {Product} from "./Entity/Product";
import {Category} from "./Entity/Category";
import {Payment} from "./ValueObject/Payment";
import {IShopRepository} from "./IShopRepository";

export type ShopPopulation = {
    readonly title: string;
    readonly ownerId: Uuid;
    readonly payment: Payment;
    readonly categories: Category[];
    readonly products: Product[];
}

export class Shop extends BaseEntity<Uuid> implements IAggregateRoot {
    private readonly title: string;
    private readonly ownerId: Uuid;
    private readonly payment: Payment;
    private readonly categories: Category[];
    private readonly products: Product[];

    constructor(private readonly repository: IShopRepository, shop?: ShopPopulation) {
        super()
        if (shop) {
            this.title = shop.title;
            this.ownerId = shop.ownerId;
            this.payment = shop.payment;
            this.categories = shop.categories;
            this.products = shop.products;
        }
    }

    create(title: string, ownerId: Uuid): Promise<Uuid> {
        //validation
        return null;
    }
}