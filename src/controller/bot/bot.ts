import Bot from 'node-telegram-bot-api';

export function run(params: any) {
    const bot = new Bot(params.token, { polling: true });

    bot.on('message', async (msg) => {
        await bot.sendMessage(msg.chat.id, "Hello");
    });

    return bot;
}

