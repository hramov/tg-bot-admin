import express from "express";
import {startController} from "../controllers/bot/start";
import {stopController} from "../controllers/bot/stop";
import {removeController} from "../controllers/bot/remove";
import {checkController} from "../controllers/bot/check";
import {newController} from "../controllers/bot/new";

export function botRouter(router: express.Router): express.Router {
    router.get('/new', newController);
    router.get('/start', startController);
    router.get('/stop', stopController);
    router.get('/remove', removeController);
    router.get('/check', checkController);
    return router;
}