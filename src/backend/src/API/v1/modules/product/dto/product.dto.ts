import {IsArray, IsJSON, IsNumber, IsString, IsOptional} from "class-validator";

export class ProductDto {

    @IsOptional()
    @IsString()
    public id: string;

    @IsString()
    public title: string;

    @IsString()
    public description: string;

    @IsArray()
    public images: string[];

    @IsNumber()
    public quantity: number;

    @IsNumber()
    public price: number;

    @IsString()
    public currency: string;

    @IsOptional()
    @IsJSON()
    public custom_fields: Record<string, string | number>;

    @IsString()
    public shop_id: string;

    @IsString()
    public category_id: string;
}