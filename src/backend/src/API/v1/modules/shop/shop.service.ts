import {Injectable} from '@nestjs/common';
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ShopSearchFilter} from "../../common/filters/shop/search.filter";
import {CreateShopDto} from "./dto/create-shop.dto";
import {Fetch} from "../../../../Infrastructure/Fetch/Fetch";
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import {ControllerResponse} from "../../../../../../shared/backendToController";
import {FetchError} from "../../../../Infrastructure/Fetch/Error";
import {telegramBotTokenRegexp, telegramUsernameRegexp} from "./shop.constants";
import {TelegramError} from "../../../../Core/Context/Shop/Error/Telegram.error";
import {InjectRepository} from "@nestjs/typeorm";
import {Repository} from "typeorm";
import { ShopEntity } from "../../common/persistent/entity/shop/shop.entity";
import {UserDto} from "../user/dto/user.dto";

@Injectable()
export class ShopService {

    constructor(@InjectRepository(ShopEntity) private readonly shopRepository: Repository<ShopEntity>) {}

    async get(filters: ShopSearchFilter) {
        return this.shopRepository.find();
    }

    async getById(shopId: Uuid) {
        return this.shopRepository.findOneBy({id: shopId.toString()});
    }

    async getByOwnerId(ownerTgName: string) {
        return this.shopRepository.findOneBy({owner_tg_name: ownerTgName});
    }

    async save(dto: CreateShopDto, user: UserDto): Promise<Uuid | Error> {
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
            return new TelegramError(data.message);
        }

        const shop = this.shopRepository.create();
        shop.owner_id = user.id;
        shop.owner_tg_name = dto.owner_tg_name;
        shop.local_shop_name = dto.local_shop_name;
        shop.bot_token = dto.bot_token;

        const result = await this.shopRepository.save(shop);
        return new Uuid(result.id);
    }

    async delete(shopId: Uuid) {
        await this.shopRepository.delete({ id: shopId.toString()});
        return shopId;
    }
}
