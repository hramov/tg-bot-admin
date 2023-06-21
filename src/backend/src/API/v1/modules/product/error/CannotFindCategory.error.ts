export class CannotFindCategoryError extends Error {
    constructor() {
        super('Cannot find category');
    }
}