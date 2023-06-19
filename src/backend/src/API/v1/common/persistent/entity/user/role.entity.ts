import {Column, Entity} from "typeorm";
import {BaseEntity} from "../entity";

@Entity({
    name: 'role',
    schema: 'user'
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