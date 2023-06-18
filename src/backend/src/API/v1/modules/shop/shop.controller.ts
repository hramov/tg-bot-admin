import {
    Body,
    Controller, Delete, Get, Param, Post, Put, Query
} from '@nestjs/common';
import {ShopService} from "./shop.service";
import {ApiBearerAuth, ApiOperation, ApiResponse, ApiTags} from "@nestjs/swagger";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ShopDto} from "./dto/shop.dto";
import {ShopSearchFilter} from "../../common/filters/shop/search.filter";
import {CreateShopDto} from "./dto/create-shop.dto";
import {Public} from "../user/public.decorator";
import {DatabaseError} from "../../../../Core/Error/Database.error";
import {checkError} from "../../error/CheckError";

@Controller('shop')
export class ShopController {
    constructor(private readonly shopService: ShopService) {}

    @ApiTags('Shop')
    @ApiBearerAuth()
    @Get('/')
    @ApiOperation({
        summary: 'Get shop list'
    })
    @ApiResponse({
        status: 200,
    })
    async get(@Query() query: string) {
        const filters = new ShopSearchFilter(query);
        return this.shopService.get(filters);
    }

    @ApiTags('Shop')
    @ApiBearerAuth()
    @Get('/owner/:owner_id')
    @ApiOperation({
        summary: 'Get shop by owner id'
    })
    @ApiResponse({
        status: 200,
    })
    async getByOwnerId(@Param('owner_id') ownerId: Uuid) {
        return this.shopService.getByOwnerId(ownerId);
    }

    @ApiTags('Shop')
    @ApiBearerAuth()
    @Get('/:id')
    @ApiOperation({
        summary: 'Get shop by id'
    })
    @ApiResponse({
        status: 200,
    })
    async getById(@Param('id') shopId: Uuid) {
        return this.shopService.getById(shopId);
    }

    @Public()
    @ApiTags('Shop')
    @ApiBearerAuth()
    @Post('/')
    @ApiOperation({
        summary: 'Create new shop',
    })
    @ApiResponse({
        status: 201,
    })
    async create(@Body() dto: CreateShopDto) {
        const data = await this.shopService.create(dto);
        if (data instanceof Error) {
            return checkError(data);
        }
        return data;
    }

    @ApiTags('Shop')
    @ApiBearerAuth()
    @Put('/:id')
    @ApiOperation({
        summary: 'Update existing shop'
    })
    @ApiResponse({
        status: 200,
    })
    async update(@Body() dto: ShopDto, @Param('id') shopId: Uuid) {
        return this.shopService.update(dto, shopId);
    }

    @ApiTags('Shop')
    @ApiBearerAuth()
    @Delete('/:id')
    @ApiOperation({
        summary: 'Delete shop'
    })
    @ApiResponse({
        status: 200,
    })
    async delete(@Param('id') shopId: Uuid) {
        return this.shopService.delete(shopId);
    }
}
