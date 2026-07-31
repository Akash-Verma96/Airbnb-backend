import { CreationOptional, DataTypes, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize from "./sequelize";


class Room extends Model<InferAttributes<Room>, InferCreationAttributes<Room>>{
    declare id: CreationOptional<number>;
    declare hotel_id: number;
    declare doa: Date;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare booking_id?: number;
    declare price: number;
    declare room_no: number
}

Room.init(
    {
       id: {
            type: DataTypes.INTEGER,
            autoIncrement: true,
            primaryKey: true,
            allowNull: false,
       },
       hotel_id: {
        type: DataTypes.INTEGER,
        allowNull: false,
        references: {
            model: 'hotels',
            key: 'id'
        },
        onUpdate: 'CASCADE',
        onDelete: 'CASCADE'
       },
       doa: {
        type: DataTypes.DATE,
        allowNull: false
       },
       createdAt: {
        type: DataTypes.DATE,
       },
       updatedAt: {
        type: DataTypes.DATE,
       },
       booking_id: {
        type: DataTypes.INTEGER,
        defaultValue: null
       },
       price: {
        type: DataTypes.FLOAT,
        allowNull: false
       },
       room_no: {
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