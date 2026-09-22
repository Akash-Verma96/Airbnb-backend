'use strict';

import { QueryInterface } from "sequelize";

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE hotels
      ADD COLUMN price INT NOT NULL,
      ADD COLUMN room_type ENUM('Single', 'Double', 'family', 'Deluxe', 'Suite') NOT NULL
    `);
  },

  async down (queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE hotels
      DROP COLUMN price,
      DROP COLUMN room_type
    `);
  }
};
