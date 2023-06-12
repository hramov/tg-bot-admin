import {
    Body,
    Controller, Get, Param, Post, Put, Query
} from '@nestjs/common';
import {ShopService} from "./shop.service";
import {ApiBearerAuth, ApiOperation, ApiResponse, ApiTags} from "@nestjs/swagger";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ShopDto} from "./dto/shop.dto";
import {ShopSearchFilter} from "../../common/filters/shop/search.filter";

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

    @ApiTags('Shop')
    @ApiBearerAuth()
    @Post('/')
    @ApiOperation({
        summary: 'Create new shop'
    })
    @ApiResponse({
        status: 200,
    })
    async create(@Body() dto: ShopDto) {
        return this.shopService.create(dto);
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
    @Get('/:id')
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
