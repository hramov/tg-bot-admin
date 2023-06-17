import express from "express";
import {homeController} from "../controllers/ui/home";

export function uiRouter(router: express.Router): express.Router {
    router.get('/', homeController);
    return router;
}