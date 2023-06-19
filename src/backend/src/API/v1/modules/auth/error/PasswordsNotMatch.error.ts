export class PasswordsDontMatchError extends Error {
    constructor() {
        super('Provided passwords don\'t match');
    }
}