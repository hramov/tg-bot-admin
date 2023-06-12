import {IEventPublisher} from "./IEventPublisher";
import {IEventBus} from "../../../Core/IEventBus";
import {ILogger} from "../../../Core/ICore";
import {BaseEvent} from "../../../Core/BaseEvent";

export class EventPublisher implements IEventPublisher {

    constructor(private readonly logger: ILogger, private readonly eventBus: IEventBus) {}

    publish<U, T extends BaseEvent<U>>(eventName: string, customEvent: T): Promise<void> {
        this.logger.log(`Got new event: ${customEvent.type} with data: ${JSON.stringify(customEvent)}`, 'EventPublisher', {
            method: 'publish'
        });

        this.eventBus.emit(eventName, customEvent);
        return;
    }

}