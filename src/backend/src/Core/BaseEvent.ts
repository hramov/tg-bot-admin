import {Ip} from "../Shared/src/ValueObject/Objects/Ip";

export class BaseEvent<T> {
    public aggregateId: string;
    public dateCreated: Date;
    public userId: string;
    public userIp: Ip;
    public revision: number;
    public data: T;
    public type: string;
}