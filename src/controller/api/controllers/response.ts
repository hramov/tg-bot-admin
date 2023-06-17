import express from "express";
import {ControllerResponse} from "../../../shared/backendToController";

export function sendResponse(res: express.Response, data: ControllerResponse) {
    res.statusCode = 200;
    return res.json(data);
}

export function sendError(res: express.Response, status: number, data: ControllerResponse) {
    res.statusCode = status;
    return res.json(data);
}