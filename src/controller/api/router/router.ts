import express from "express";
import {botRouter} from "./bot";
import {uiRouter} from "./ui";

export function register(router: express.Router) {
    router.use('/bot', botRouter(router));
    router.use('/ui', uiRouter(router));
}