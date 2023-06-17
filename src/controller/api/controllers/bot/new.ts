import express from "express";
import {v4} from "uuid";
import {run} from "../../../bot/bot";
import {botStore} from "../../../store/bot.store";

export async function newController(req: express.Request, res: express.Response) {
    const token = req.query.token;

    if (!token) {
        res.statusCode = 400;
        res.json({
            status: "Error",
            message: 'No token provided'
        });
        return
    }

    if (botStore.state.tokens.has(token.toString())) {
        res.statusCode = 400;
        res.json({
            status: "Error",
            message: 'The bot is already running',
        });
        return
    }

    const id = v4();
    const bot = run({token: token});
    botStore.state.tokens.set(id, token.toString());

    botStore.mutations.set(id, bot);
    res.json({
        status: "OK",
        id: id,
    });
}