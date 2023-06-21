import {Column, Entity, JoinColumn, JoinTable, ManyToMany, ManyToOne, OneToMany, OneToOne} from "typeorm";
import {BaseEntity} from "./entity";
import {ShopEntity} from "./shop.entity";
import {RoleEntity} from "./role.entity";

@Entity({
    name: 'user',
})
export class UserEntity extends BaseEntity {

    @Column({
        name: 'tg_name',
    })
    public tg_name: string;

    @Column({
        name: 'password',
    })
    public password: string;

    @ManyToOne(() => RoleEntity, role => role.id)
    @JoinColumn({
        name: 'role',
    })
    public role: RoleEntity; // admin, shop owner, shop manager

    @OneToMany(() => ShopEntity, shop => shop.owner)
    shop: ShopEntity;
}