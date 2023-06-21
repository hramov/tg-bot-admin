export const telegramUsernameRegexp = new RegExp('.*?\\B\\w{3,64}\\b.*');
export const telegramBotTokenRegexp = new RegExp(/[0-9]{9}:[a-zA-Z0-9_-]{35}/gm);