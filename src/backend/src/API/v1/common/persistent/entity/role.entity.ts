import {Column, Entity, JoinColumn, ManyToOne, OneToMany} from "typeorm";
import {BaseEntity} from "./entity";
import {UserEntity} from "./user.entity";

@Entity({
    name: 'role',
})
export class RoleEntity extends BaseEntity {

    @Column({
        name: 'title',
        type: 'varchar',
        unique: true,
    })
    public title: string;

    @Column({
        name: 'code',
        type: 'integer',
        unique: true,
    })
    public code: number;
}