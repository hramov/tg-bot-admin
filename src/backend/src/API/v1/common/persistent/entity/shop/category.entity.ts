import {Column, Entity, OneToMany} from "typeorm";
import {BaseEntity} from "../entity";
import {ProductEntity} from "./product.entity";

@Entity({
    name: 'category',
    schema: 'shop'
})
export class CategoryEntity extends BaseEntity {
    @Column({
        unique: true
    })
    title: string;

    @OneToMany(() => ProductEntity, product => product.category_id)
    products: ProductEntity[];

}