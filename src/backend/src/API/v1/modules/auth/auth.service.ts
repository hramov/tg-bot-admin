import { Injectable } from '@nestjs/common';
import {RegisterDto} from "./dto/register.dto";
import {LoginDto} from "./dto/login.dto";

@Injectable()
export class AuthService {

    register(dto: RegisterDto) {
        return Promise.resolve(undefined);
    }

    login(dto: LoginDto) {
        return Promise.resolve(undefined);
    }
}
