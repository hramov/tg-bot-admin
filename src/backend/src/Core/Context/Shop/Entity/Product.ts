import {BaseEntity} from "../../../../Shared/src/BaseEntity";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {Price} from "../ValueObject/Price";

export class Product extends BaseEntity<Uuid> {
    private readonly shopId: Uuid;
    private readonly categoryId: Uuid;

    private readonly title: string;
    private readonly description: string;
    private readonly images: string[];
    private readonly amount: number;
    private readonly price: Price;
}