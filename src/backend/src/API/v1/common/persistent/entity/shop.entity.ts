import {Column, Entity, JoinColumn, ManyToOne, OneToMany, OneToOne, PrimaryGeneratedColumn} from "typeorm";
import {BaseEntity} from "./entity";
import {ProductEntity} from "./product.entity";
import {UserEntity} from "./user.entity";
import {OrderEntity} from "./order.entity";

@Entity({
    name: 'shop',
})
export class ShopEntity extends BaseEntity {
    @Column()
    local_shop_name: string;

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

    @OneToMany(() => ProductEntity, product => product.shop)
    products: ProductEntity[];

    @OneToMany(() => OrderEntity, order => order.shop)
    orders: OrderEntity[];

    @ManyToOne(() => UserEntity, user => user.shop)
    owner: UserEntity;
}