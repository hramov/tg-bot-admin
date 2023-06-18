import {IsString} from "class-validator";
import {ApiProperty} from "@nestjs/swagger";

export class CreateShopDto {

    @ApiProperty({
        type: 'string',
        description: 'Telegram bot token'
    })
    @IsString()
    bot_token: string;

    @ApiProperty({
        type: 'string',
        description: 'Local shop name',
    })
    @IsString()
    local_shop_name: string;

    @ApiProperty({
        type: 'string',
        description: 'Owner telegram username'
    })
    @IsString()
    owner_tg_name: string;
}