

import { DataTypes, QueryInterface } from "sequelize";


module.exports = {
  async up (queryInterface:QueryInterface) {
   await queryInterface.addColumn(`hotels`, `deletedAt`, {
      type: DataTypes.DATE,
      allowNull: true,
      defaultValue: null,
    });
  },

  async down (queryInterface:QueryInterface) {
   await queryInterface.removeColumn(`hotels`, `deletedAt`);
  }
};
