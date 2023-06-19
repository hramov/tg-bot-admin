import {Inject, Injectable} from '@nestjs/common';
import {RegisterDto} from "./dto/register.dto";
import {LoginDto} from "./dto/login.dto";
import {InjectRepository} from "@nestjs/typeorm";
import {UserEntity} from "../../common/persistent/entity/user/user.entity";
import {Repository} from "typeorm";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {PasswordsDontMatchError} from "./error/PasswordsNotMatch.error";
import {telegramUsernameRegexp} from "../shop/shop.constants";
import {WrongTelegramUsernameError} from "../../error/WrongTelegramUsername.error";
import { genSalt, hash } from 'bcrypt';
import {RoleEntity} from "../../common/persistent/entity/user/role.entity";
import {CannotFindRoleError} from "./error/CannotFindRole.error";
import { compare} from "bcrypt";
import {LoginOrPasswordIncorrectError} from "./error/LoginOrPasswordIncorrect.error";
import {JwtService} from "@nestjs/jwt";

@Injectable()
export class AuthService {

    constructor(
        @InjectRepository(UserEntity) private readonly userRepository: Repository<UserEntity>,
        @InjectRepository(RoleEntity) private readonly roleRepository: Repository<RoleEntity>,
        private readonly jwtService: JwtService
    ) {}

    async register(dto: RegisterDto): Promise<Uuid | Error> {

        if (dto.password !== dto.confirm_password) {
            return new PasswordsDontMatchError();
        }

        if (!dto.tg_name.match(telegramUsernameRegexp)) {
            return new WrongTelegramUsernameError();
        }

        const user = this.userRepository.create();
        user.tg_name = dto.tg_name;

        const role = await this.roleRepository.findOneBy({code: dto.role_code});
        if (!role || !role.id) {
            return new CannotFindRoleError();
        }

        user.role_id = role.id;

        const salt = await genSalt(5)
        user.password = await hash(dto.password, salt);

        const result = await this.userRepository.save(user);
        return new Uuid(result.id);
    }

    async login(dto: LoginDto) {
        const user = await this.userRepository.findOneBy({tg_name: dto.tg_name});

        const valid = await compare(dto.password, user.password);

        if (!valid) {
            return new LoginOrPasswordIncorrectError();
        }

        const role = await this.roleRepository.findOneBy({id: user.role_id });
        if (!role || !role.id) {
            return new CannotFindRoleError();
        }

        const payload = { id: user.id, tg_name: user.tg_name, role_code: role.code };
        return {
            access_token: await this.jwtService.signAsync(payload),
        };
    }
}
