import {IsBoolean, IsEmail, IsStrongPassword} from "class-validator";

export class LoginDto {
    @IsEmail()
    email: string;

    @IsStrongPassword()
    password: string;

    @IsBoolean()
    remember: boolean;
}