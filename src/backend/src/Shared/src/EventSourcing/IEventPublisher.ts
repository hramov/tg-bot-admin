import {BaseEvent} from "../../../Core/BaseEvent";

export interface IEventPublisher {
    publish<U, T extends BaseEvent<U>>(eventName: string, event: T): Promise<void>
}