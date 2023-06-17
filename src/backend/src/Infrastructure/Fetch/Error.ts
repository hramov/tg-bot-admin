export class FetchError extends Error {
    constructor(url: string, err: Error) {
        super(`Cannot fetch data from ${url}: ${err.message}`)
    }
}