'use strict';

import { QueryInterface } from "sequelize";


/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      CREATE TABLE IF NOT EXISTS room_categories (
        id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
        hotel_id INT NOT NULL,
        room_type ENUM('Single', 'Double', 'family', 'Deluxe', 'Suite') NOT NULL,
        room_no INT NOT NULL,
        price INT NOT NULL,

        created_at DATETIME,
        updated_at DATETIME,
        deleted_at DATETIME
      )
    `)
  },

  async down (queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      DROP TABLE IF EXISTS room_categories
    `)
  }
};
