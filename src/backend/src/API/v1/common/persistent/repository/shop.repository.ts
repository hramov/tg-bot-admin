import {ILogger} from "../../../../../Core/ICore";
import {IDatabaseConnection} from "../../../../../Infrastructure/Persistent/Storage/IDatabaseConnection";
import {IPostgresQueryOptions} from "../../../../../Infrastructure/Persistent/Storage/Postgres/IPostgresQueryOptions";
import {IShopRepository} from "../../../../../Core/Context/Shop/IShopRepository";
import {ShopSearchFilter} from "../../filters/shop/search.filter";
import {ShopConstructor} from "../../../../../Core/Context/Shop/Shop";
import {Uuid} from "../../../../../Shared/src/ValueObject/Objects/Uuid";
import {DatabaseError} from "../../../../../Core/Error/Database.error";
import {CreateShopDto} from "../../../modules/shop/dto/create-shop.dto";

export class ShopRepository implements IShopRepository {
    constructor(private readonly logger: ILogger, private readonly storage: IDatabaseConnection<IPostgresQueryOptions>) {}

    async create(dto: CreateShopDto): Promise<Uuid | DatabaseError> {
        this.logger.log('Creating the shop', 'ShopRepository', {
            method: 'create'
        });
        const sql = `
            insert into shop.shop (shop_id, local_shop_name, owner_tg_name, bot_token)
            values (uuid_generate_v4(), $1, $2, $3)
            returning shop_id
        `;

        const params = [dto.local_shop_name, dto.owner_tg_name, dto.bot_token];
        const res = await this.storage.queryOne<Uuid>(sql, params);
        if (res instanceof DatabaseError) {
            this.logger.warn(res.message, 'ShopRepository');
        }
        return res;
    }

    getByFilters(filters: ShopSearchFilter): Promise<ShopConstructor[] | DatabaseError> {
        return Promise.resolve(undefined);
    }

    getById(id: Uuid): Promise<ShopConstructor | DatabaseError> {
        return Promise.resolve(undefined);
    }

    getByOwnerId(ownerId: Uuid): Promise<ShopConstructor | DatabaseError> {
        return Promise.resolve(undefined);
    }


}