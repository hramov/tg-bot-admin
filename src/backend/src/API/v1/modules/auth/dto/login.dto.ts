import {IsString, IsStrongPassword} from "class-validator";

export class LoginDto {
    @IsString()
    tg_name: string;

    @IsStrongPassword()
    password: string;
}