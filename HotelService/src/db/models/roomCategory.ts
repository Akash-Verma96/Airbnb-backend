import { CreationOptional, DataTypes, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize from "./sequelize";

export enum RoomType {
  SINGLE = 'SINGLE',
  DOUBLE = 'DOUBLE',
  FAMILY = 'FAMILY',
  DELUXE = 'DELUXE',
  SUITE = 'SUITE',
}

class RoomCategory extends Model<InferAttributes<RoomCategory>, InferCreationAttributes<RoomCategory>>{
    declare id: CreationOptional<number>;
    declare hotelId: number;
    declare roomType: RoomType;
    declare roomNo: number;
    declare price: number;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare deletedAt: CreationOptional<Date> | null;
}

RoomCategory.init({
    id: {
        type: DataTypes.INTEGER,
        autoIncrement: true,
        primaryKey: true,
        allowNull: false
    },
    hotelId: {
        type: DataTypes.INTEGER,
        allowNull: false,
    },
    roomType: {
      type: 'ENUM',
      values: [...Object.values(RoomType)],
    },
    roomNo: {
        type: DataTypes.INTEGER,
        allowNull: false
    },
    price: {
        type: DataTypes.INTEGER,
        allowNull: false
    },
    createdAt: {
        type: DataTypes.DATE,
       },
    updatedAt: {
        type: DataTypes.DATE,
    },
    deletedAt: {
        type: DataTypes.DATE,
        defaultValue: null,
    }
},{
    tableName: 'room_categories',
    sequelize: sequelize,
    underscored: true,
    timestamps: true
})

export default RoomCategory;