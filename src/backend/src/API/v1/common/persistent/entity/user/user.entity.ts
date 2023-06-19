import {Column, Entity, JoinColumn, JoinTable, ManyToMany, ManyToOne, PrimaryGeneratedColumn} from "typeorm";
import {BaseEntity} from "../entity";
import {ShopEntity} from "../shop/shop.entity";
import {RoleEntity} from "./role.entity";

@Entity({
    name: 'user',
    schema: 'user'
})
export class UserEntity extends BaseEntity {

    @Column({
        name: 'tg_name',
    })
    public tg_name: string;

    @ManyToOne(() => RoleEntity, role => role.code)
    @JoinColumn({
        name: 'role_id'
    })
    public role_id: string; // admin, shop owner, shop manager

    @ManyToMany(() => ShopEntity, { eager: true })
    @JoinTable({
        name: "user_shop",
        schema: 'user',
        joinColumn: {
            name: "user_id",
            referencedColumnName: "id"
        },
        inverseJoinColumn: {
            name: "shop_id",
            referencedColumnName: "id"
        }
    })
    shops: ShopEntity[];
}