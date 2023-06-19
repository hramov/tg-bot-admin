import {Column, Generated, PrimaryGeneratedColumn} from "typeorm";

export class BaseEntity {
    @PrimaryGeneratedColumn('uuid')
    @Generated('uuid')
    public id: string;

    @Column({
        default: new Date(),
    })
    public date_created: Date;

    @Column({
        default: null
    })
    public last_updated: Date;
}