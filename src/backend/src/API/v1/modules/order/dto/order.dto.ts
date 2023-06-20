import {CoreProduct} from "../../../../../Core/Context/Shop/Entity/Product";
import {IsArray, IsString, IsUUID} from "class-validator";
import {ApiProperty} from "@nestjs/swagger";

export class OrderDto {

    @ApiProperty({
        type: 'uuid',
        description: 'Shop id to serve order',
    })
    @IsUUID()
    shop_id: string;

    @ApiProperty({
        type: 'uuid',
        description: 'Manager to serve order',
    })
    @IsUUID()
    manager_id: string;

    @ApiProperty({
        type: 'string',
        description: 'Telegram username of customer',
    })
    @IsString()
    customer_tg_name: string;

    @ApiProperty({
        type: 'string',
        description: 'Where to deliver the order',
    })
    @IsString()
    delivery_address: string;

    @ApiProperty({
        type: 'array',
        description: 'Products in order',
    })
    @IsArray()
    products: CoreProduct[];
}