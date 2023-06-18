import {ILogger} from "../../../../../Core/ICore";
import {IDatabaseConnection} from "../../../../../Infrastructure/Persistent/Storage/IDatabaseConnection";
import {IPostgresQueryOptions} from "../../../../../Infrastructure/Persistent/Storage/Postgres/IPostgresQueryOptions";
import {IShopRepository} from "../../../../../Core/Context/Shop/IShopRepository";
import {ShopSearchFilter} from "../../filters/shop/search.filter";
import {Shop} from "../../../../../Core/Context/Shop/Shop";
import {Uuid} from "../../../../../Shared/src/ValueObject/Objects/Uuid";
import {DatabaseError} from "../../../../../Core/Error/Database.error";

export class ShopRepository implements IShopRepository {
    constructor(private readonly logger: ILogger, private readonly storage: IDatabaseConnection<IPostgresQueryOptions>) {}

    getByFilters(filters: ShopSearchFilter): Promise<Shop[]> {
        return Promise.resolve([]);
    }

    getById(id: Uuid): Promise<Shop> {
        return Promise.resolve(undefined);
    }

    getByOwnerId(ownerId: Uuid): Promise<Shop | DatabaseError> {
        return Promise.resolve(undefined);
    }


}