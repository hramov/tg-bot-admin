import express from "express";
import {startController} from "../controllers/start";
import {stopController} from "../controllers/stop";
import {removeController} from "../controllers/remove";
import {checkController} from "../controllers/check";
import {newController} from "../controllers/new";

export function botRouter(router: express.Router): express.Router {
    router.get('/new', newController);
    router.get('/start', startController);
    router.get('/stop', stopController);
    router.get('/remove', removeController);
    router.get('/check', checkController);
    return router;
}