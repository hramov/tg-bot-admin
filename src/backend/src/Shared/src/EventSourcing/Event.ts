import Emitter from "events";

export class Event {
    private readonly event: Emitter;
    constructor() {
        this.event = new Emitter()
    }
}