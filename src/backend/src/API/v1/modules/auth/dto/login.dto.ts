import {IsString, IsStrongPassword} from "class-validator";
import {ApiProperty} from "@nestjs/swagger";

export class LoginDto {

    @ApiProperty({
        type: 'string',
        description: 'Telegram username',
        examples: [
            '@therealhramov'
        ]
    })
    @IsString()
    tg_name: string;

    @ApiProperty({
        type: 'string',
        description: 'Password',
        examples: [
            'qWerty!23456'
        ]
    })
    @IsStrongPassword()
    password: string;
}