import { QueryInterface } from "sequelize";

module.exports = {
  async up(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE Hotels
      ADD COLUMN address VARCHAR(255) NOT NULL,
      ADD COLUMN rating_count INT DEFAULT 0
    `);
  },

  async down(queryInterface: QueryInterface) {
    await queryInterface.sequelize.query(`
      ALTER TABLE Hotels
      DROP COLUMN address,
      DROP COLUMN rating_count
    `);
  }
};