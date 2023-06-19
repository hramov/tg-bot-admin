export class CannotFindRoleError extends Error {
    constructor() {
        super('Cannot find role');
    }
}