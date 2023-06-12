import {ShopSearchFilter} from "../../../API/v1/common/filters/shop/search.filter";
import {Shop} from "./Shop";
import {Uuid} from "../../../Shared/src/ValueObject/Objects/Uuid";
import {DatabaseError} from "../../Error/Database.error";

export interface IShopRepository {
    getByFilters(filters: ShopSearchFilter): Promise<Shop[] | DatabaseError>;
    getById(id: Uuid): Promise<Shop | DatabaseError>;
}