import {Column, Entity, JoinColumn, JoinTable, ManyToMany, ManyToOne} from "typeorm";
import {BaseEntity} from "./entity";
import {Uuid} from "../../../../../Shared/src/ValueObject/Objects/Uuid";
import {TelegramName} from "../../../../../Core/Context/Shop/ValueObject/TelegramId";
import {Address} from "../../../../../Core/Context/Shop/ValueObject/Address";
import {CoreProduct} from "../../../../../Core/Context/Shop/Entity/Product";
import {OrderStatus} from "../../../../../Core/Context/Shop/Entity/Order";
import {ShopEntity} from "./shop.entity";
import {ProductEntity} from "./product.entity";
import {UserEntity} from "./user.entity";

@Entity({
    name: 'order',
})
export class OrderEntity extends BaseEntity {

    @ManyToOne(() => ShopEntity, shop => shop.orders)
    @JoinColumn({
        name: 'shop'
    })
    shop: Uuid;

    @ManyToOne(() => UserEntity, manager => manager.shop)
    @JoinColumn({
        name: 'manager'
    })
    manager: UserEntity;

    @Column({
        type: 'varchar',
        name: 'customer_tg_name'
    })
    customer_tg_name: TelegramName;

    @Column({
        type: 'varchar',
        name: 'courier_name'
    })
    courier_name: string;

    @Column({
        type: 'varchar',
        name: 'delivery_address'
    })
    delivery_address: Address;

    @Column({
        type: 'varchar',
        name: 'order_tracking_link'
    })
    order_tracking_link: string;

    @ManyToMany(() => ProductEntity, { eager: true })
    @JoinTable({
        name: "order_product",
        joinColumn: {
            name: "product_id",
            referencedColumnName: "id"
        },
        inverseJoinColumn: {
            name: "order_id",
            referencedColumnName: "id"
        }
    })
    products: CoreProduct[];

    @Column({
        type: 'varchar',
        name: 'status'
    })
    status: OrderStatus;

    @Column({
        type: 'timestamp',
        name: 'order_time'
    })
    order_time: Date;

    @Column({
        type: 'timestamp',
        name: 'expected_time'
    })
    expected_time: Date;

    @Column({
        type: 'timestamp',
        name: 'sending_time'
    })
    sending_time: Date;

    @Column({
        type: 'timestamp',
        name: 'delivery_time'
    })
    delivery_time: Date;

}