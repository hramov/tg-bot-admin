import {Column, Entity, JoinColumn, OneToMany, OneToOne, PrimaryGeneratedColumn} from "typeorm";
import {BaseEntity} from "../entity";
import {ProductEntity} from "./product.entity";
import {UserEntity} from "../user/user.entity";

@Entity({
    name: 'shop',
    schema: 'shop'
})
export class ShopEntity extends BaseEntity {
    @Column()
    local_shop_name: string;

    @OneToOne(() => UserEntity, user => user.id)
    @Column({
        name: 'owner_id',
        type: 'uuid',
    })
    owner_id: string;

    @Column({
        name: 'owner_tg_name'
    })
    owner_tg_name: string;

    @Column({
        name: 'bot_token'
    })
    bot_token: string;

    @OneToMany(() => ProductEntity, product => product.shop_id)
    products: ProductEntity[];
}