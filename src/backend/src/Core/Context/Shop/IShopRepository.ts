import {ShopSearchFilter} from "../../../API/v1/common/filters/shop/search.filter";
import {Shop, ShopConstructor} from "./Shop";
import {Uuid} from "../../../Shared/src/ValueObject/Objects/Uuid";
import {DatabaseError} from "../../Error/Database.error";
import {CreateShopDto} from "../../../API/v1/modules/shop/dto/create-shop.dto";

export interface IShopRepository {
    getByFilters(filters: ShopSearchFilter): Promise<ShopConstructor[] | DatabaseError>;
    getById(id: Uuid): Promise<ShopConstructor | DatabaseError>;
    getByOwnerId(ownerId: Uuid): Promise< ShopConstructor | DatabaseError>;
    create(dto: CreateShopDto): Promise<Uuid | DatabaseError>;
}