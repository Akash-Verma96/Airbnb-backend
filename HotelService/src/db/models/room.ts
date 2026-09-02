import { CreationOptional, DataTypes, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize from "./sequelize";
import RoomCategory from "./roomCategory";


class Room extends Model<InferAttributes<Room>, InferCreationAttributes<Room>>{
    declare id: CreationOptional<number>;
    declare hotelId: number;
    declare roomCategoryId: number;
    declare dateOfAvailability: Date;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare deletedAt: CreationOptional<Date> | null;
    declare booking_id?: number;
    declare price: number;
    declare roomNo: number;
}

Room.init(
    {
       id: {
            type: DataTypes.INTEGER,
            autoIncrement: true,
            primaryKey: true,
            allowNull: false,
       },
       roomCategoryId: {
            type: 'INTEGER',
            allowNull: false,
            references: {
                model: RoomCategory,
                key: 'id',
            },
        },
       hotelId: {
        type: DataTypes.INTEGER,
        allowNull: false,
        references: {
            model: 'hotels',
            key: 'id'
        },
        onUpdate: 'CASCADE',
        onDelete: 'CASCADE'
       },
       dateOfAvailability: {
        type: DataTypes.DATE,
        allowNull: false
       },
       createdAt: {
        type: DataTypes.DATE,
       },
       updatedAt: {
        type: DataTypes.DATE,
       },
       deletedAt: {
        type: 'DATE',
        defaultValue: null,
        },
       booking_id: {
        type: DataTypes.INTEGER,
        defaultValue: null
       },
       price: {
        type: DataTypes.FLOAT,
        allowNull: false
       },
       roomNo: {
        type: DataTypes.INTEGER,
        allowNull: false
       }
    },{
        tableName: 'rooms',
        sequelize: sequelize,
        underscored: true,
        timestamps: true
    }
)

export default Room;