import {
    Controller, Get, HttpCode, Post,
} from '@nestjs/common';
import {UserService} from "./user.service";
import {ApiBearerAuth, ApiOperation, ApiResponse, ApiTags} from "@nestjs/swagger";
import {User} from "../auth/user.decorator";
import {UserDto} from "./dto/user.dto";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {checkError} from "../../error/CheckError";

@Controller('user')
export class UserController {
    constructor(private readonly userService: UserService) {}

    @ApiTags('User')
    @Get('/info')
    @HttpCode(200)
    @ApiBearerAuth()
    @ApiOperation({
        summary: 'Get user info via token'
    })
    @ApiResponse({
        status: 200,
    })
    async info(@User() user: UserDto) {
        const result = await this.userService.getById(new Uuid(user.id));
        if (result instanceof Error) {
            checkError(result);
        }
        return result;
    }

}
