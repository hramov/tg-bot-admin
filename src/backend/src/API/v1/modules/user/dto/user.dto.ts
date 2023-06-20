import {IsNumber, IsString, IsUUID} from "class-validator";
import {ApiProperty} from "@nestjs/swagger";

export class UserDto {

    @ApiProperty({
        type: 'uuid',
        description: 'User id',
    })
    @IsUUID()
    public id: string;

    @ApiProperty({
        type: 'string',
        description: 'User telegram username',
    })
    @IsString()
    public tg_name: string;

    @ApiProperty({
        type: 'number',
        description: 'User role code',
    })
    @IsNumber()
    public role_code: string;
}