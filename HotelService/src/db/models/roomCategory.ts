import { CreationOptional, DataTypes, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize from "./sequelize";

enum RoomType {
  SINGLE = 'SINGLE',
  DOUBLE = 'DOUBLE',
  FAMILY = 'FAMILY',
  DELUXE = 'DELUXE',
  SUITE = 'SUITE',
}

class RoomCategory extends Model<InferAttributes<RoomCategory>, InferCreationAttributes<RoomCategory>>{
    declare id: CreationOptional<number>;
    declare hotel_id: number;
    declare roomType: RoomType;
    declare roomNo: number;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare deletedAt: CreationOptional<Date>;
}

RoomCategory.init({
    id: {
        type: DataTypes.INTEGER,
        autoIncrement: true,
        primaryKey: true,
        allowNull: false
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
    roomType: {
      type: 'ENUM',
      values: [...Object.values(RoomType)],
    },
    roomNo: {
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
        type: DataTypes.DATE
    }
},{
    tableName: 'room_categories',
    sequelize: sequelize,
    underscored: true,
    timestamps: true
})

export default RoomCategory;