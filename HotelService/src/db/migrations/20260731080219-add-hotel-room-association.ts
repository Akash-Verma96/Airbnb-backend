'use strict';

import { QueryInterface } from "sequelize";

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface: QueryInterface) {
    await queryInterface.addConstraint('rooms', {
      fields: ['hotel_id'],
      type: 'foreign key',
      name: 'custom_fkey_room_hotels',
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
      'hotels',
      'custom_fkey_room_hotels'
    );
  }
};
