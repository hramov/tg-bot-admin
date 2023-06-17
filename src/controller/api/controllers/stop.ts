import express from "express";
import {botStore} from "../../store/bot.store";

export async function stopController(req: express.Request, res: express.Response) {
    const id = req.query.id;
    if (!id) {
        res.statusCode = 400;
        res.send('Wrong query');
        return
    }

    const bot = botStore.mutations.get(id.toString());
    if (!bot) {
        res.statusCode = 500;
        res.send('Cannot get bot');
        return;
    }

    await bot.stopPolling();
    res.json({
        status: "OK"
    });
}