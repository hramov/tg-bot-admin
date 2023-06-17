import {BaseEntity} from "../../../Shared/src/BaseEntity";
import {Uuid} from "../../../Shared/src/ValueObject/Objects/Uuid";
import {IAggregateRoot} from "../../../Shared/src/IAggregateRoot";
import {Product} from "./Entity/Product";
import {Category} from "./Entity/Category";
import {Payment} from "./ValueObject/Payment";
import {IShopRepository} from "./IShopRepository";
import {Fetch} from "../../../Infrastructure/Fetch/Fetch";
import {FetchError} from "../../../Infrastructure/Fetch/Error";

export type ShopConstructor = {
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

    constructor(private readonly repository: IShopRepository, shop?: ShopConstructor) {
        super()
        if (shop) {
            this.title = shop.title;
            this.ownerId = shop.ownerId;
            this.payment = shop.payment;
            this.categories = shop.categories;
            this.products = shop.products;
        }
    }

    private mapToObj(): ShopConstructor {
        return {
            title: this.title,
            ownerId: this.ownerId,
            payment: this.payment,
            categories: this.categories,
            products: this.products,
        }
    }

    public async create(ownerId: Uuid, title: string, imgUrl: string): Promise<Uuid> {
        return null;
    }

    public async linkBot(shopId: Uuid, token: string): Promise<Uuid> {
        if (await this.checkIsTokenCorrect(token)) {
            // create bot instance in controller, but not started
        }
        return null;
    }

    public async startBot(botId: Uuid): Promise<Uuid> {
        // make api call to controller to route /start?id=botId
        return null;
    }

    private async checkIsTokenCorrect(token: string): Promise<boolean> {
        const data = await Fetch.get<{ status: string }>('http://localhost:3002/check?token=' + token);
        if (data instanceof FetchError) {
            return false;
        }
        return data.status === 'valid';
    }
}