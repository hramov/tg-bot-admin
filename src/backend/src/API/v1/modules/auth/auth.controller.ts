import {
    Body,
    Controller, Post,
} from '@nestjs/common';
import {ApiBearerAuth, ApiOperation, ApiResponse, ApiTags} from "@nestjs/swagger";
import {AuthService} from "./auth.service";
import {LoginDto} from "./dto/login.dto";
import {RegisterDto} from "./dto/register.dto";

@Controller('auth')
export class AuthController {
    constructor(private readonly authService: AuthService) {}

    @ApiTags('Auth')
    @ApiBearerAuth()
    @Post('/login')
    @ApiOperation({
        summary: 'Login'
    })
    @ApiResponse({
        status: 200,
    })
    async login(@Body() dto: LoginDto) {
        return this.authService.login(dto);
    }

    @ApiTags('Auth')
    @ApiBearerAuth()
    @Post('/register')
    @ApiOperation({
        summary: 'Register'
    })
    @ApiResponse({
        status: 200,
    })
    async register(@Body() dto: RegisterDto) {
        return this.authService.register(dto);
    }
}
