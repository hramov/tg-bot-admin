import {DatabaseError} from "../../../Core/Error/Database.error";
import {BadRequestException} from "@nestjs/common";

export function checkError(error: Error) {
    if (error instanceof DatabaseError) {
        return checkDatabaseError(error);
    }
}

function checkDatabaseError(error: DatabaseError) {
    const msg = error.message;
    if (msg.startsWith('duplicate key value violates unique constraint')) {
        const wrongFieldArray = msg.split('"')[1].split('_')
        const wrongField = wrongFieldArray.filter((_, index) => index > 0 && index < wrongFieldArray.length - 1).join('_');
        return new BadRequestException('Field ' + wrongField + ' is already exists');
    }
}