import {
    Body,
    Controller, Delete, Get, Param, Post, Put, Query,
} from '@nestjs/common';
import {ProductService} from "./product.service";
import {ApiBearerAuth, ApiOperation, ApiResponse, ApiTags} from "@nestjs/swagger";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ProductDto} from "./dto/product.dto";
import {ProductSearchFilter} from "../../common/filters/product/search.filter";

@Controller('product')
export class ProductController {
    constructor(private readonly productService: ProductService) {}

    @ApiTags('Product')
    @ApiBearerAuth()
    @Get('/')
    @ApiOperation({
        summary: 'Get all products by filters'
    })
    @ApiResponse({
        status: 200,
    })
    async get(@Query() query: string) {
        const filters = new ProductSearchFilter(query);
        return this.productService.get(filters);
    }

    @ApiTags('Product')
    @ApiBearerAuth()
    @Get('/:id')
    @ApiOperation({
        summary: 'Get product by id'
    })
    @ApiResponse({
        status: 200,
    })
    async getById(@Param('id') productId: Uuid) {
        return this.productService.getById(productId);
    }

    @ApiTags('Product')
    @ApiBearerAuth()
    @Post('/')
    @ApiOperation({
        summary: 'Create new product'
    })
    @ApiResponse({
        status: 200,
    })
    async create(@Body() dto: ProductDto) {
        return this.productService.create(dto);
    }

    @ApiTags('Product')
    @ApiBearerAuth()
    @Put('/:id')
    @ApiOperation({
        summary: 'Update product'
    })
    @ApiResponse({
        status: 200,
    })
    async update(@Body() dto: ProductDto, @Param('id') productId: Uuid) {
        return this.productService.update(productId);
    }

    @ApiTags('Product')
    @ApiBearerAuth()
    @Delete('/:id')
    @ApiOperation({
        summary: 'Delete product'
    })
    @ApiResponse({
        status: 200,
    })
    async delete(@Param('id') productId: Uuid) {
        return this.productService.delete(productId);
    }
}
