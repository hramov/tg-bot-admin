import {BaseEntity} from "../../../../Shared/src/BaseEntity";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";

export class Category extends BaseEntity<Uuid> {
    private readonly title: string;
}