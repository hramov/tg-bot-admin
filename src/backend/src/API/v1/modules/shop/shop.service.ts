import {Inject, Injectable} from '@nestjs/common';
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ShopDto} from "./dto/shop.dto";
import {ShopSearchFilter} from "../../common/filters/shop/search.filter";
import {IShopRepository} from "../../../../Core/Context/Shop/IShopRepository";
import {SHOP_REPOSITORY} from "../../common/persistent/repository/repository.constants";
import {CreateShopDto} from "./dto/create-shop.dto";
import {Fetch} from "../../../../Infrastructure/Fetch/Fetch";
import {ControllerResponse} from "../../../../../../shared/backendToController";
import {FetchError} from "../../../../Infrastructure/Fetch/Error";
import {telegramBotTokenRegexp, telegramUsernameRegexp} from "./shop.constants";

@Injectable()
export class ShopService {

    constructor(@Inject(SHOP_REPOSITORY) private readonly shopRepository: IShopRepository) {}

    get(filters: ShopSearchFilter) {
        return this.shopRepository.getByFilters(filters);
    }

    getById(shopId: Uuid) {
        return this.shopRepository.getById(shopId);
    }

    getByOwnerId(shopId: Uuid) {
        return this.shopRepository.getByOwnerId(shopId);
    }

    async create(dto: CreateShopDto): Promise<Uuid | Error> {
        if (dto.owner_tg_name.match(telegramUsernameRegexp) === null) {
            return new Error('Wrong telegram username format');
        }

        if (dto.bot_token.match(telegramBotTokenRegexp) === null) {
            return new Error('Wrong bot token format');
        }

        const data = await Fetch.get<ControllerResponse>(process.env.CONTROLLER_URL + '/check?token=' + dto.bot_token);
        if (data instanceof FetchError) {
            return new Error('Cannot check token');
        }

        if (data.status === false) {
            return new Error(data.message);
        }

        return this.shopRepository.create(dto);
    }

    update(dto: ShopDto, shopId: Uuid) {
        return Promise.resolve(undefined);
    }

    delete(shopId: Uuid) {
        return Promise.resolve(undefined);
    }
}
