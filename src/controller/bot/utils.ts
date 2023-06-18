import Bot from "node-telegram-bot-api";
import {botStore} from "../store/bot.store";

export async function removeBot(bot: Bot, id?: string) {
    await bot.stopPolling();
    await bot.removeAllListeners();
    if (id) {
        botStore.mutations.delete(id as string);
    }
}