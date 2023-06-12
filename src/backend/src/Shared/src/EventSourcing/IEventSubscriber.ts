export interface IEventSubscriber {
    subscribe<T>(event: T): Event
}