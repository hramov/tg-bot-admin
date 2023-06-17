import express from "express";
import {v4} from "uuid";
import {run} from "../../../bot/bot";
import {botStore} from "../../../store/bot.store";
import {sendError, sendResponse} from "../response";

export async function newController(req: express.Request, res: express.Response) {
    const token = req.query.token;

    if (!token) {
        sendError(res, 400, {
            status: false,
            message: 'No token provided'
        });
        return
    }

    if (botStore.state.tokens.has(token.toString())) {
        sendError(res, 400, {
            status: false,
            message: 'The bot is already running',
        });
        return
    }

    const id = v4();
    const bot = run({token: token});
    botStore.state.tokens.set(id, token.toString());

    botStore.mutations.set(id, bot);
    sendResponse(res, {
        status: true,
        message: id,
    });
}