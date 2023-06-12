import {DatabaseError} from "../../../Core/Error/Database.error";

export interface IDatabaseConnection<U> {
    query: <T>(
        sql: string,
        values?: any[],
        options?: U,
    ) => Promise<T[] | DatabaseError>

    queryOne: <T>(
        sql: string,
        values?: any[],
        options?: U,
    ) => Promise<T | DatabaseError>

    queryTx: <T>(
        sql: string[],
        values?: any[][],
        options?: U,
    ) => Promise<any | DatabaseError>
}