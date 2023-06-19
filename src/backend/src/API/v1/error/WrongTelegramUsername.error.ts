export class WrongTelegramUsernameError extends Error {
    constructor() {
        super('Wrong telegram username format');
    }
}