'use strict';

import { QueryInterface } from "sequelize";

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface: QueryInterface) {
    await queryInterface.addConstraint('room_categories', {
      fields: ['hotel_id'],
      type: 'foreign key',
      name: 'custom_fkey_room_category_hotel',
      references: { //Required field
        table: 'hotels',
        field: 'id'
      },
      onDelete: 'cascade',
      onUpdate: 'cascade'
    });
  },

  async down (queryInterface: QueryInterface) {
    await queryInterface.removeConstraint(
      'room_categories',
      'custom_fkey_room_category_hotel'
    );
  }
};
