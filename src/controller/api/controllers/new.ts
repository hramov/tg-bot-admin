import express from "express";
import {v4} from "uuid";
import {run} from "../../bot/bot";
import {botStore} from "../../store/bot.store";

export async function newController(req: express.Request, res: express.Response) {
    const token = req.query.token;
    const id = v4();
    botStore.mutations.set(id, run({token: token}))
    res.json({
        status: "OK",
        id: id,
    });
}