import {
    Body,
    Controller, Delete, Get, Param, Post, Query,
} from '@nestjs/common';
import {ApiBearerAuth, ApiOperation, ApiResponse, ApiTags} from "@nestjs/swagger";
import {OrderService} from "./order.service";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {OrderDto} from "./dto/order.dto";
import {OrderSearchFilter} from "../../common/filters/order/search.filter";
import {checkError} from "../../error/CheckError";

@Controller('order')
export class OrderController {
    constructor(private readonly orderService: OrderService) {}

    @ApiTags('Order')
    @ApiBearerAuth()
    @Get('/')
    @ApiOperation({
        summary: 'Get all orders by filters'
    })
    @ApiResponse({
        status: 200,
    })
    async get(@Query() query: string) {
        const filters = new OrderSearchFilter(query);
        const result = await this.orderService.get(filters);
        if (result instanceof Error) {
            checkError(result);
        }
        return result;
    }

    @ApiTags('Order')
    @ApiBearerAuth()
    @Get('/:id')
    @ApiOperation({
        summary: 'Get order by id'
    })
    @ApiResponse({
        status: 200,
    })
    async getById(@Param('id') orderId: Uuid) {
        return this.orderService.getById(orderId);
    }

    @ApiTags('Order')
    @ApiBearerAuth()
    @Post('/')
    @ApiOperation({
        summary: 'Create new order'
    })
    @ApiResponse({
        status: 200,
    })
    async create(@Body() dto: OrderDto) {
        return this.orderService.create(dto);
    }

    @ApiTags('Order')
    @ApiBearerAuth()
    @Delete('/:id')
    @ApiOperation({
        summary: 'Cancel order'
    })
    @ApiResponse({
        status: 200,
    })
    async cancel(@Param('id') orderId: Uuid) {
        return this.orderService.cancel(orderId);
    }
}
