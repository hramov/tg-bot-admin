import {
    Body,
    Controller, Delete, Get, Param, Post, Put, Query,
} from '@nestjs/common';
import {ProductService} from "./product.service";
import {ApiBearerAuth, ApiBody, ApiOperation, ApiResponse, ApiTags} from "@nestjs/swagger";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ProductDto} from "./dto/product.dto";
import {ProductSearchFilter} from "../../common/filters/product/search.filter";
import {checkError} from "../../error/CheckError";

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
        type: ProductDto,
        isArray: true,
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
        type: ProductDto,
    })
    async getById(@Param('id') productId: Uuid) {
        return this.productService.getById(productId);
    }

    @ApiTags('Product')
    @ApiBearerAuth()
    @Post('/')
    @ApiOperation({
        summary: 'Create new product',
    })
    @ApiBody({
        type: ProductDto
    })
    @ApiResponse({
        status: 200,
    })
    async create(@Body() dto: ProductDto) {
        const result = await this.productService.save(dto);
        if (result instanceof Error) {
            checkError(result);
        }
        return result;
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
    async update(@Body() dto: ProductDto) {
        const result = await this.productService.save(dto);
        if (result instanceof Error) {
            checkError(result);
        }
        return result;
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
