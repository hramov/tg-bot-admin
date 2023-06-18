import {IsString} from "class-validator";
import {ApiProperty} from "@nestjs/swagger";

export class CreateShopDto {

    @ApiProperty({
        type: 'string',
        description: 'Telegram bot token',
        examples: [
            '5945818399:AAHk2WbGLvhD3PqbG0Vg3onzX_vdb_s9SNs'
        ]
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
        description: 'Owner telegram username',
        examples: [
            '@therealhramov'
        ]
    })
    @IsString()
    owner_tg_name: string;
}