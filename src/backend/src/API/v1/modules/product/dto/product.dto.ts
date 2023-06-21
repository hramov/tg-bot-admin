import {IsArray, IsJSON, IsNumber, IsString, IsOptional, IsUUID} from "class-validator";
import {ApiProperty} from "@nestjs/swagger";

export class ProductDto {

    @ApiProperty({
        type: 'uuid',
        description: 'If you need to update existing product, pass the id',
    })
    @IsOptional()
    @IsUUID()
    public id: string;

    @ApiProperty({
        type: 'string',
        description: 'Product title',
    })
    @IsString()
    public title: string;

    @ApiProperty({
        type: 'string',
        description: 'Product description',
    })
    @IsString()
    public description: string;

    @ApiProperty({
        type: 'array',
        description: 'Array of images',
    })
    @IsArray()
    public images: string[];

    @ApiProperty({
        type: 'number',
        description: 'Product quantity',
    })
    @IsNumber()
    public quantity: number;

    @ApiProperty({
        type: 'number',
        description: 'Product price',
    })
    @IsNumber()
    public price: number;

    @ApiProperty({
        type: 'string',
        description: 'Price currency',
    })
    @IsString()
    public currency: string;

    @ApiProperty({
        type: 'json',
        description: 'Custom fields to be displayed',
    })
    @IsJSON()
    public custom_fields: Record<string, string | number>;

    @ApiProperty({
        type: 'uuid',
        description: 'What shop product belongs',
    })
    @IsUUID()
    public shop_id: string;

    @ApiProperty({
        type: 'uuid',
        description: 'What category product belongs',
    })
    @IsUUID()
    public category_id: string;
}