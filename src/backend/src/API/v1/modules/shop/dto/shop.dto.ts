import { IsString, IsUrl} from "class-validator";

export class ShopDto {
    @IsString()
    title: string;

    @IsString()
    description: string;

    @IsUrl()
    imgUrl: string;
}