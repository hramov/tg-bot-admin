import {ShopService} from "../../../src/API/v1/modules/shop/shop.service";
import {IShopRepository} from "../../../src/Core/Context/Shop/IShopRepository";
import {Uuid} from "../../../src/Shared/src/ValueObject/Objects/Uuid";
import {DatabaseError} from "../../../src/Core/Error/Database.error";
import {ShopConstructor} from "../../../src/Core/Context/Shop/Shop";
import {ShopSearchFilter} from "../../../src/API/v1/common/filters/shop/search.filter";
import {CreateShopDto} from "../../../src/API/v1/modules/shop/dto/create-shop.dto";
import { v4 } from 'uuid';

describe('Shop service', () => {
    let shopRepository: IShopRepository = null;

    beforeAll(() => {
        shopRepository = {
            async getByOwnerId(ownerId: Uuid): Promise<ShopConstructor | DatabaseError> {
                return null;
            },
            async getById(id: Uuid): Promise<ShopConstructor | DatabaseError> {
                return null;
            },
            async getByFilters(filters: ShopSearchFilter): Promise<ShopConstructor[] | DatabaseError> {
                return null;
            },
            async create(dto: CreateShopDto): Promise<Uuid | DatabaseError> {
                if (dto.local_shop_name === '123') {
                    return new DatabaseError('Database Error');
                }
                return new Uuid(v4().toString());
            }
        }
    });

    describe('Create Shop', () => {
        it('Wrong owner telegram username', async () => {
            const shopService = new ShopService(shopRepository);

            const dto: CreateShopDto = {
                local_shop_name: 'Test Shop',
                owner_tg_name: '123',
                bot_token: '',
            }

            const result = await shopService.create(dto);
            expect(result).toBeInstanceOf(Error);
            expect((result as Error).message).toBe('Wrong telegram username format')
        });

        it('Wrong telegram bot token', async () => {
            const shopService = new ShopService(shopRepository);

            const dto: CreateShopDto = {
                local_shop_name: 'Test Shop',
                owner_tg_name: '@therealhramov',
                bot_token: 'wrongBotToken',
            }

            const result = await shopService.create(dto);
            expect(result).toBeInstanceOf(Error);
            expect((result as Error).message).toBe('Wrong bot token format')
        });

        it('Database Error', async () => {
            const shopService = new ShopService(shopRepository);

            const dto: CreateShopDto = {
                local_shop_name: '123',
                owner_tg_name: '@therealhramov',
                bot_token: '5945818399:AAHk2WbGLvhD3PqbG0Vg3onzX_vdb_s9SNs',
            }

            const result = await shopService.create(dto);
            expect(result).toBeInstanceOf(DatabaseError);
        });

        it('Success', async () => {
            const shopService = new ShopService(shopRepository);

            const dto: CreateShopDto = {
                local_shop_name: 'Test Shop',
                owner_tg_name: '@therealhramov',
                bot_token: '5945818399:AAHk2WbGLvhD3PqbG0Vg3onzX_vdb_s9SNs',
            }

            const result = await shopService.create(dto);
            expect(result).toBeInstanceOf(Uuid);
        });
    })
})