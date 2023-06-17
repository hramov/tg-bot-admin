import express from "express";
import {botStore} from "../../../store/bot.store";
import {sendError, sendResponse} from "../response";

export async function startController(req: express.Request, res: express.Response) {
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

    await bot.startPolling();

    sendResponse(res, {
        status: true,
        message: '',
    });
}