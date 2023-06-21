import {ApiProperty} from "@nestjs/swagger";
import {IsNumber, IsString, IsStrongPassword} from "class-validator";

export class RegisterDto {

    @ApiProperty({
        type: 'string',
        description: 'Telegram username',
        examples: ['@therealhramov']
    })
    @IsString()
    public tg_name: string;

    @ApiProperty({
        type: 'number',
        description: 'Role code',
        examples: [10]
    })
    @IsNumber()
    public role_code: number;

    @ApiProperty({
        type: 'string',
        description: 'Plain password, should be strong',
        examples: ['qWerty!23456']
    })
    @IsString()
    @IsStrongPassword()
    public password: string;

    @ApiProperty({
        type: 'string',
        description: 'Same plain password',
        examples: ['qWerty!23456']
    })
    @IsString()
    @IsStrongPassword()
    public confirm_password: string;
}