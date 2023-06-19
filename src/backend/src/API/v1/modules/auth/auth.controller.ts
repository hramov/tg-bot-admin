import {
    Body,
    Controller, Post,
} from '@nestjs/common';
import {ApiOperation, ApiResponse, ApiTags} from "@nestjs/swagger";
import {AuthService} from "./auth.service";
import {LoginDto} from "./dto/login.dto";
import {RegisterDto} from "./dto/register.dto";
import {checkError} from "../../error/CheckError";
import {Public} from "./public.decorator";

@Controller('auth')
export class AuthController {
    constructor(private readonly authService: AuthService) {}

    @ApiTags('Auth')
    @Post('/login')
    @Public()
    @ApiOperation({
        summary: 'Login'
    })
    @ApiResponse({
        status: 200,
    })
    async login(@Body() dto: LoginDto) {
        const result = await this.authService.login(dto);
        if (result instanceof Error) {
            checkError(result)
        }
        return result;
    }

    @ApiTags('Auth')
    @Post('/register')
    @Public()
    @ApiOperation({
        summary: 'Register'
    })
    @ApiResponse({
        status: 200,
    })
    async register(@Body() dto: RegisterDto) {
        const result = await this.authService.register(dto);
        if (result instanceof Error) {
            checkError(result)
        }
        return result;
    }
}
