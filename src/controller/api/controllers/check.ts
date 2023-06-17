import express from "express";
import {run} from "../../bot/bot";
import {removeBot} from "../../bot/utils";

export async function checkController(req: express.Request, res: express.Response) {
    const token = req.query.token;
    let reason = '';

    const checkPromise = await new Promise((resolve, reject) => {
        const bot = run({token: token});

        bot.on('polling_error', (err) => {
            reason = err.message;
            removeBot(bot)
            resolve(false)
        });

        setTimeout(() => resolve(true), 2000);
    });

    if (checkPromise) {
        res.statusCode = 200;
        return res.json({
            status: "Valid",
        });
    }

    res.statusCode = 400;
    res.json({
        status: "Not valid",
        reason: reason,
    });
}