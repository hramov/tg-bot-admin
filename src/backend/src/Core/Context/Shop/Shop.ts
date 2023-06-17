import {BaseEntity} from "../../../Shared/src/BaseEntity";
import {Uuid} from "../../../Shared/src/ValueObject/Objects/Uuid";
import {IAggregateRoot} from "../../../Shared/src/IAggregateRoot";
import {Product} from "./Entity/Product";
import {Category} from "./Entity/Category";
import {Payment} from "./ValueObject/Payment";
import {IShopRepository} from "./IShopRepository";
import {Fetch} from "../../../Infrastructure/Fetch/Fetch";
import {FetchError} from "../../../Infrastructure/Fetch/Error";
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import {ControllerResponse} from "../../../../../shared/backendToController";

export type ShopConstructor = {
    readonly title: string;
    readonly ownerId: Uuid;
    readonly token: string;
    readonly botId: string;
    readonly payment: Payment;
    readonly categories: Category[];
    readonly products: Product[];
}

export class Shop extends BaseEntity<Uuid> implements IAggregateRoot {
    private readonly title: string;
    private readonly ownerId: Uuid;
    private token: string;
    private botId: Uuid;
    private payment: Payment;
    private readonly categories: Category[];
    private readonly products: Product[];

    constructor(private readonly repository: IShopRepository, shop?: ShopConstructor) {
        super()
        if (shop) {
            this.title = shop.title;
            this.ownerId = shop.ownerId;
            this.token = shop.token;
            this.botId = new Uuid(shop.botId);
            this.payment = shop.payment;
            this.categories = shop.categories;
            this.products = shop.products;
        }
    }

    private mapToObj(): ShopConstructor {
        return {
            title: this.title,
            ownerId: this.ownerId,
            token: this.token,
            botId: this.botId.toString(),
            payment: this.payment,
            categories: this.categories,
            products: this.products,
        }
    }

    private static async checkIsTokenCorrect(token: string): Promise<boolean> {
        const data = await Fetch.get<ControllerResponse>(process.env.CONTROLLER_URL + '/check?token=' + token);
        if (data instanceof FetchError) {
            return false;
        }
        return data.status === true;
    }

    public async create(ownerId: Uuid, title: string, imgUrl: string): Promise<Uuid> {
        return null;
    }

    public async update(ownerId: Uuid, title: string, imgUrl: string): Promise<Uuid> {
        return null;
    }

    public async linkBot(token: string): Promise<Uuid> {
        if (await Shop.checkIsTokenCorrect(token)) {
            this.token = token;

            const botId = await this.createBot();
            if (botId) {
                this.botId = new Uuid(botId);
                return this.botId;
            }
        }
        return null;
    }

    private async createBot(): Promise<string> {
        const data = await Fetch.get<ControllerResponse>(process.env.CONTROLLER_URL + '/new?token=' + this.token);
        if (data instanceof FetchError) {
            return null;
        }
        return data.message;
    }

    public async startBot(): Promise<Uuid> {
        const data = await Fetch.get<ControllerResponse>(process.env.CONTROLLER_URL + '/api/bot/start?id=' + this.botId);
        if (data instanceof FetchError) {
            return null;
        }
        return new Uuid(data.message);
    }

    public async stopBot(): Promise<Uuid> {
        const data = await Fetch.get<ControllerResponse>(process.env.CONTROLLER_URL + '/api/bot/stop?id=' + this.botId);
        if (data instanceof FetchError) {
            return null;
        }
        return new Uuid(data.message);
    }

    public async removeBot(): Promise<Uuid> {
        const data = await Fetch.get<ControllerResponse>(process.env.CONTROLLER_URL + '/api/bot/remove?id=' + this.botId);
        if (data instanceof FetchError) {
            return null;
        }
        return new Uuid(data.message);
    }

    public async delete(ownerId: Uuid, title: string, imgUrl: string): Promise<Uuid> {
        return null;
    }
}