import {Column, Entity, JoinColumn, ManyToOne} from "typeorm";
import {BaseEntity} from "../entity";
import {ShopEntity} from "./shop.entity";
import {CategoryEntity} from "./category.entity";

@Entity({
    name: 'product',
    schema: 'shop'
})
export class ProductEntity extends BaseEntity {
    @Column()
    public title: string;

    @Column()
    public description: string;

    @Column({
        type: 'jsonb'
    })
    public images: string[];

    @Column()
    public quantity: number;

    @Column()
    public price: number;

    @Column()
    public currency: string;

    @Column({
        type: 'jsonb',
        name: 'custom_fields'
    })
    public customFields: Record<string, string | number>;

    @ManyToOne(() => ShopEntity, shop => shop.products)
    @JoinColumn({
        name: 'shop_id'
    })
    public shop_id: ShopEntity;

    @ManyToOne(() => CategoryEntity, category => category.products)
    @JoinColumn({
        name: 'category_id'
    })
    public category_id: string;
}