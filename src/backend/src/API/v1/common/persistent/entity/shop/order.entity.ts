import {Column, Entity, JoinColumn, JoinTable, ManyToMany, ManyToOne} from "typeorm";
import {BaseEntity} from "../entity";
import {Uuid} from "../../../../../../Shared/src/ValueObject/Objects/Uuid";
import {TelegramName} from "../../../../../../Core/Context/Shop/ValueObject/TelegramId";
import {Address} from "../../../../../../Core/Context/Shop/ValueObject/Address";
import {CoreProduct} from "../../../../../../Core/Context/Shop/Entity/Product";
import {OrderStatus} from "../../../../../../Core/Context/Shop/Entity/Order";
import {ShopEntity} from "./shop.entity";
import {ProductEntity} from "./product.entity";
import {UserEntity} from "../user/user.entity";

@Entity({
    name: 'order',
    schema: 'shop'
})
export class OrderEntity extends BaseEntity {

    @ManyToOne(() => ShopEntity, shop => shop.id)
    @JoinColumn({
        name: 'shop_id'
    })
    private readonly shop_id: Uuid;

    @ManyToOne(() => UserEntity, manager => manager.id)
    @JoinColumn({
        name: 'manager_id'
    })
    private readonly manager_id: Uuid;

    @Column({
        type: 'varchar',
        name: 'customer_tg_name'
    })
    private readonly customer_tg_name: TelegramName;

    @Column({
        type: 'varchar',
        name: 'courier_name'
    })
    private readonly courier_name: string;

    @Column({
        type: 'varchar',
        name: 'delivery_address'
    })
    private readonly delivery_address: Address;

    @Column({
        type: 'varchar',
        name: 'order_tracking_link'
    })
    private readonly order_tracking_link: string;

    @ManyToMany(() => ProductEntity, { eager: true })
    @JoinTable({
        name: "order_product",
        schema: 'shop',
        joinColumn: {
            name: "product_id",
            referencedColumnName: "id"
        },
        inverseJoinColumn: {
            name: "order_id",
            referencedColumnName: "id"
        }
    })
    private readonly products: CoreProduct[];

    @Column({
        type: 'varchar',
        name: 'status'
    })
    private readonly status: OrderStatus;

    @Column({
        type: 'timestamp',
        name: 'order_time'
    })
    private readonly order_time: Date;

    @Column({
        type: 'timestamp',
        name: 'expected_time'
    })
    private readonly expected_time: Date;

    @Column({
        type: 'timestamp',
        name: 'sending_time'
    })
    private readonly sending_time: Date;

    @Column({
        type: 'timestamp',
        name: 'delivery_time'
    })
    private readonly delivery_time: Date;

}