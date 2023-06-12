import express from 'express';
import {run} from "./bot/bot";
import type Bot from 'node-telegram-bot-api';
import { v4 } from 'uuid';

const bots = new Map<string, Bot>();

function main() {
    const app = express();

    app.get('/new', (req, res) => {
        const token = req.query.token;
        const id = v4();
        bots.set(id, run({token: token}))
        res.json({
            status: "OK",
            id: id,
        });
    });

    app.get('/start', async (req, res) => {
        const id = req.query.id;
        if (!id) {
            res.statusCode = 400;
            res.send('Wrong query');
            return
        }

        const bot = bots.get(id.toString());
        if (!bot) {
            res.statusCode = 500;
            res.send('Cannot get bot');
            return;
        }

        await bot.startPolling()

        res.json({
            status: "OK"
        })
    });

    app.get('/stop', async (req, res) => {
        const id = req.query.id;
        if (!id) {
            res.statusCode = 400;
            res.send('Wrong query');
            return
        }

        const bot = bots.get(id.toString());
        if (!bot) {
            res.statusCode = 500;
            res.send('Cannot get bot');
            return;
        }

        await bot.stopPolling();
        res.json({
            status: "OK"
        })
    });

    app.get('/remove', async (req, res) => {
        const id = req.query.id;
        if (!id) {
            res.statusCode = 400;
            res.send('Wrong query');
            return
        }

        const bot = bots.get(id.toString());
        if (!bot) {
            res.statusCode = 500;
            res.send('Cannot get bot');
            return;
        }

        await bot.stopPolling()
        await bot.removeAllListeners();
        bots.delete(id as string);

        res.json({
            status: "OK",
        })
    });

    app.listen(3002, () => {
        console.log('Controller started on port 3002');
    })
}

main();