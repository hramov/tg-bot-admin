import express from "express";
import {botStore} from "../../../store/bot.store";
import {removeBot} from "../../../bot/utils";
import {sendError, sendResponse} from "../response";

export async function removeController(req: express.Request, res: express.Response) {
    const id = req.query.id;
    if (!id) {
       sendError(res, 400, {
            status: false,
            message: 'Wrong query'
        });
        return
    }

    const bot = botStore.mutations.get(id.toString());
    if (!bot) {
        sendError(res, 400, {
            status: false,
            message: 'Cannot find bot'
        });
        return;
    }

    await removeBot(bot, id.toString());

    botStore.state.tokens.delete(id.toString());

    sendResponse(res, {
        status: true,
        message: '',
    });
}