import Bot from "node-telegram-bot-api";

export const botStore = {
    state: {
        bots: new Map<string, Bot>(),
    },
    mutations: {
        set: (key: string, value: Bot) => {
            botStore.state.bots.set(key, value);
        },
        get: (key: string) => {
            return botStore.state.bots.get(key);
        },
        delete: (key: string) => {
            botStore.state.bots.delete(key);
        },
        has: (key: string) => {
            return botStore.state.bots.has(key);
        },
    }
}

