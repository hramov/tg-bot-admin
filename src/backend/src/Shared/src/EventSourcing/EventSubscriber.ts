import {IEventSubscriber} from "./IEventSubscriber";

export class EventSubscriber implements IEventSubscriber{
    subscribe<T>(event: T): Event {
        return undefined;
    }

}