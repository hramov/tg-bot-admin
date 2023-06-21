import {
    Body,
    Controller, Delete, Get, HttpCode, Param, Post, Put, Query
} from '@nestjs/common';
import {ShopService} from "./shop.service";
import {ApiBearerAuth, ApiOperation, ApiResponse, ApiTags} from "@nestjs/swagger";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ShopSearchFilter} from "../../common/filters/shop/search.filter";
import {ShopDto} from "./dto/shop.dto";
import {checkError} from "../../error/CheckError";
import {Roles} from "../auth/roles.decorator";
import {Role} from "../auth/role.enum";
import {User} from "../auth/user.decorator";
import {UserDto} from "../user/dto/user.dto";

@Controller('shop')
export class ShopController {
    constructor(private readonly shopService: ShopService) {}

    @ApiTags('Shop')
    @ApiBearerAuth()
    @Roles(Role.Admin)
    @Get('/')
    @HttpCode(200)
    @ApiOperation({
        summary: 'Get shop list'
    })
    @ApiResponse({
        status: 200,
    })
    async get(@Query() query: string) {
        const filters = new ShopSearchFilter(query);
        const result = await this.shopService.get(filters);
        if (result instanceof Error) {
            checkError(result);
        }
        return result;
    }

    @ApiTags('Shop')
    @ApiBearerAuth()
    @Get('/my')
    @HttpCode(200)
    @Roles(Role.Admin, Role.Owner)
    @ApiOperation({
        summary: 'Get user shop'
    })
    @ApiResponse({
        status: 200,
    })
    async getByOwnerId(@User() user: UserDto) {
        const result = await this.shopService.getByOwnerId(user.id);
        if (result instanceof Error) {
            checkError(result);
        }
        return result;
    }

    @ApiTags('Shop')
    @ApiBearerAuth()
    @Get('/:id')
    @HttpCode(200)
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
    @HttpCode(201)
    @Roles(Role.Admin, Role.Owner)
    @ApiOperation({
        summary: 'Create new shop',
    })
    @ApiResponse({
        status: 201,
    })
    @ApiResponse({
        status: 400,
    })
    @ApiResponse({
        status: 500,
    })
    async create(@Body() dto: ShopDto, @User() user: UserDto) {
        const data = await this.shopService.save(dto, user);
        if (data instanceof Error) {
            checkError(data);
        }
        return data;
    }

    @ApiTags('Shop')
    @ApiBearerAuth()
    @Put('/:id')
    @HttpCode(204)
    @ApiOperation({
        summary: 'Update existing shop'
    })
    @ApiResponse({
        status: 200,
    })
    async update(@Body() dto: ShopDto, @User() user: UserDto) {
        const data = await this.shopService.save(dto, user);
        if (data instanceof Error) {
            checkError(data);
        }
        return data;
    }

    @ApiTags('Shop')
    @ApiBearerAuth()
    @Delete('/:id')
    @HttpCode(204)
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
