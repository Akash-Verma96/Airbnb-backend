'use strict';

import { QueryInterface } from "sequelize";

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface: QueryInterface) {
    await queryInterface.addConstraint('rooms', {
      fields: ['room_category_id'],
      type: 'foreign key',
      name: 'custom_fkey_room_category',
      references: { //Required field
        table: 'room_categories',
        field: 'id'
      },
      onDelete: 'cascade',
      onUpdate: 'cascade'
    });
  },

  async down (queryInterface: QueryInterface) {
    await queryInterface.removeConstraint(
      'rooms',
      'custom_fkey_room_category'
    );
  }
};
