import {ShopService} from "../../../src/API/v1/modules/shop/shop.service";
import {Uuid} from "../../../src/Shared/src/ValueObject/Objects/Uuid";
import {DatabaseError} from "../../../src/Core/Error/Database.error";
import {ShopConstructor} from "../../../src/Core/Context/Shop/Shop";
import {CreateShopDto} from "../../../src/API/v1/modules/shop/dto/create-shop.dto";
import {ShopEntity} from "../../../src/API/v1/common/persistent/entity/shop.entity";

describe('Shop service', () => {
    // let shopRepository: ShopEntity = null;
    //
    // beforeAll(() => {
    //     shopRepository = {
    //         async create(ownerId: Uuid): Promise<ShopConstructor | DatabaseError> {
    //             return null;
    //         },
    //         async save(id: Uuid): Promise<ShopConstructor | DatabaseError> {
    //             return null;
    //         },
    //     }
    // });
    //
    // describe('Create Shop', () => {
    //     it('Wrong owner telegram username', async () => {
    //         const shopService = new ShopService(shopRepository);
    //
    //         const dto: CreateShopDto = {
    //             local_shop_name: 'Test Shop',
    //             owner_tg_name: '123',
    //             bot_token: '',
    //         }
    //
    //         const result = await shopService.save(dto);
    //         expect(result).toBeInstanceOf(Error);
    //         expect((result as Error).message).toBe('Wrong telegram username format')
    //     });
    //
    //     it('Wrong telegram bot token', async () => {
    //         const shopService = new ShopService(shopRepository);
    //
    //         const dto: CreateShopDto = {
    //             local_shop_name: 'Test Shop',
    //             owner_tg_name: '@therealhramov',
    //             bot_token: 'wrongBotToken',
    //         }
    //
    //         const result = await shopService.save(dto);
    //         expect(result).toBeInstanceOf(Error);
    //         expect((result as Error).message).toBe('Wrong bot token format')
    //     });
    //
    //     it('Database Error', async () => {
    //         const shopService = new ShopService(shopRepository);
    //
    //         const dto: CreateShopDto = {
    //             local_shop_name: '123',
    //             owner_tg_name: '@therealhramov',
    //             bot_token: '5945818399:AAHk2WbGLvhD3PqbG0Vg3onzX_vdb_s9SNs',
    //         }
    //
    //         const result = await shopService.save(dto);
    //         expect(result).toBeInstanceOf(DatabaseError);
    //     });
    //
    //     it('Success', async () => {
    //         const shopService = new ShopService(shopRepository);
    //
    //         const dto: CreateShopDto = {
    //             local_shop_name: 'Test Shop',
    //             owner_tg_name: '@therealhramov',
    //             bot_token: '5945818399:AAHk2WbGLvhD3PqbG0Vg3onzX_vdb_s9SNs',
    //         }
    //
    //         const result = await shopService.save(dto);
    //         expect(result).toBeInstanceOf(Uuid);
    //     });
    // })
})