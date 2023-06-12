export class DatabaseError extends Error {
    constructor(message?: string) {
        super('Database error: ' + message ? message : '');
    }
}