import Bot from 'node-telegram-bot-api';

export function run(params: any) {
    const bot = new Bot(params.token, { polling: true });

    bot.on('message', async (msg) => {
        const chatId = msg.chat.id;
        // const url = process.env.WEB_APP_URL + `?chat_id=${chatId}#/`;
        const text = msg.text;
        const url = 'https://hramovdev.ru'

        if (text === '/start') {
            const opts = {
                reply_markup: {
                    inline_keyboard: [
                        [
                            {
                                text: 'Click here',
                                web_app: {
                                    url: url,
                                },
                            },
                        ],
                    ],
                    resize_keyboard: true,
                },
            };

            await bot.sendMessage(chatId, 'Open the shop', opts);
        }
    });

    return bot;
}

