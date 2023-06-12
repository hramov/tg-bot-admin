export class Filter {
    constructor(private readonly query: string) {}

    // parse query string (filter=value&filter2=value2) into map[field]value
    parse(): Map<string, string> {
        const map = new Map<string, string>();
        const params = this.query.split('&');
        for (const p of params) {
            const field = p.split('=')[0];
            const value = p.split('=')[1];
            map.set(field, value);
        }
        return map;
    }
}