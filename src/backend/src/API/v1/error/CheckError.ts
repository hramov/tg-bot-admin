import {DatabaseError} from "../../../Core/Error/Database.error";
import {BadRequestException, InternalServerErrorException} from "@nestjs/common";
import {TelegramError} from "../../../Core/Context/Shop/Error/Telegram.error";

export function checkError(error: Error) {
    if (error instanceof DatabaseError) {
        checkDatabaseError(error);
    }
    if (error instanceof TelegramError) {
        checkTelegramError(error);
    }
    throw new InternalServerErrorException(error.message);
}

function checkDatabaseError(error: DatabaseError) {
    const msg = error.message;
    if (msg.startsWith('duplicate key value violates unique constraint')) {
        const wrongFieldArray = msg.split('"')[1].split('_')
        const wrongField = wrongFieldArray.filter((_, index) => index > 0 && index < wrongFieldArray.length - 1).join('_');
        throw new BadRequestException('Field ' + wrongField + ' is already exists');
    }
    throw new InternalServerErrorException(msg);
}

function checkTelegramError(error: TelegramError) {
    throw new BadRequestException(error.message);
}