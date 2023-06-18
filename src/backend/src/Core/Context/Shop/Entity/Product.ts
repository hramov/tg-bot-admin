import {BaseEntity} from "../../../../Shared/src/BaseEntity";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {Price} from "../ValueObject/Price";

export interface CoreProduct {
    product_id: Uuid;
    price: Price;
    quantity: number;
}

type CustomFields = {
    [key: string]: string | number
};

export class Product extends BaseEntity<Uuid> implements CoreProduct {
    private readonly _shop_id: Uuid;
    private readonly _category_id: Uuid;
    private readonly _product_id: Uuid;

    private readonly _title: string;
    private readonly _description: string;
    private readonly _images: string[];
    private readonly _quantity: number;
    private readonly _price: Price;
    private readonly _customFields: CustomFields;

    get title() { return this._title; }
    get price() { return this._price; }
    get quantity() { return this._quantity; }
    get product_id() { return this._product_id; }

}