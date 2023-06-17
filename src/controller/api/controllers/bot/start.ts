import express from "express";
import {botStore} from "../../../store/bot.store";

export async function startController(req: express.Request, res: express.Response) {
    const id = req.query.id;
    if (!id) {
        res.statusCode = 400;
        res.send('Wrong query');
        return
    }

    const bot = botStore.mutations.get(id.toString());
    if (!bot) {
        res.statusCode = 400;
        res.send('Cannot find bot');
        return;
    }

    await bot.startPolling();

    res.json({
        status: "OK"
    });
}