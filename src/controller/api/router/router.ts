import express from "express";
import {botRouter} from "./bot";

export function register(router: express.Router) {
    router.use('/bot', botRouter(router));
}