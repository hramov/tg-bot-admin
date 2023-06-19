export class LoginOrPasswordIncorrectError extends Error {
    constructor() {
        super('Login or password is incorrect');
    }
}