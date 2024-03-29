import express from "express";
import {run} from "../../../bot/bot";
import {removeBot} from "../../../bot/utils";
import {sendError, sendResponse} from "../response";

export async function checkController(req: express.Request, res: express.Response) {
    const token = req.query.token;
    let reason = '';

    console.log('Check token ' + token)

    const checkPromise = await new Promise((resolve, reject) => {
        const bot = run({token: token});

        bot.on('polling_error', (err) => {
            reason = err.message;
            removeBot(bot);
            resolve(false);
        });

        setTimeout(() => resolve(true), 10000);
    });

    if (checkPromise) {
        console.log('Token ' + token + ' is valid');
        return sendResponse(res, { status: true, message: '' })
    }

    console.log('Token ' + token + ' is not valid: ' + reason);
    return sendResponse(res, { status: false, message: reason })
}