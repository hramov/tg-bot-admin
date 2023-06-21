import {ForbiddenException, Injectable} from '@nestjs/common';
import {RegisterDto} from "./dto/register.dto";
import {LoginDto} from "./dto/login.dto";
import {PasswordsDontMatchError} from "./error/PasswordsNotMatch.error";
import {telegramUsernameRegexp} from "../shop/shop.constants";
import {WrongTelegramUsernameError} from "../../error/WrongTelegramUsername.error";
import { compare} from "bcrypt";
import {LoginOrPasswordIncorrectError} from "./error/LoginOrPasswordIncorrect.error";
import {JwtService} from "@nestjs/jwt";
import {UserService} from "../user/user.service";
import {UserAlreadyExistsError} from "./error/UserAlreadyExists.error";
import {secret, secretRefresh} from "./auth.constants";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";

@Injectable()
export class AuthService {

    constructor(
        private readonly jwtService: JwtService,
        private readonly userService: UserService
    ) {
    }

    async register(dto: RegisterDto): Promise<{ access_token: string, refresh_token: string } | Error> {

        const existing = await this.userService.getByTgName(dto.tg_name);

        if (existing) {
            return new UserAlreadyExistsError();
        }

        if (dto.password !== dto.confirm_password) {
            return new PasswordsDontMatchError();
        }

        if (!dto.tg_name.match(telegramUsernameRegexp)) {
            return new WrongTelegramUsernameError();
        }

        const user = await this.userService.create(dto);

        if (user instanceof Error) {
            return user;
        }

        const payload = {id: user.id, tg_name: user.tg_name, role_code: user.role.code};

        return this.getTokens(payload);
    }

    async login(dto: LoginDto) {
        const user = await this.userService.getByTgName(dto.tg_name);

        if (!user) {
            return new LoginOrPasswordIncorrectError();
        }

        const valid = await compare(dto.password, user.password);

        if (!valid) {
            return new LoginOrPasswordIncorrectError();
        }

        const payload = {id: user.id, tg_name: user.tg_name, role_code: user.role.code};
        return this.getTokens(payload);
    }

    private async getTokens(payload: { id: string, tg_name: string, role_code: number}) {
        const [access_token, refresh_token] = await Promise.all([
            this.jwtService.signAsync(
                payload,
                {
                    secret: secret,
                    expiresIn: '15m',
                },
            ),
            this.jwtService.signAsync(
                payload,
                {
                    secret: secretRefresh,
                    expiresIn: '7d',
                },
            ),
        ]);
        return {
            access_token,
            refresh_token,
        };

    }

    async refreshTokens(refreshToken: string) {
        const refreshData = await this.jwtService.verifyAsync(refreshToken, {
            secret: secretRefresh,
        });
        const user = await this.userService.getById(new Uuid(refreshData.id))
        if (!user) throw new ForbiddenException('Access Denied');

        return this.getTokens({ id: user.id, tg_name: user.tg_name, role_code: user.role.code});
    }

}
