'use strict';

import { QueryInterface, DataTypes } from "sequelize";

/** @type {import('sequelize-cli').Migration} */
export default {
  async up (queryInterface: QueryInterface) {
    await queryInterface.addColumn('hotels', 'host_id', {
      type: DataTypes.INTEGER,
      allowNull: false,
      defaultValue: 1
    });
  },

  async down (queryInterface: QueryInterface) {
    await queryInterface.removeColumn('hotels', 'host_id');
  }
};
