const { Sequelize, DataTypes } = require('sequelize');
const sequelize = new Sequelize('sqlite::memory:');

const User = sequelize.define('user', {
  firstName: DataTypes.TEXT,
  lastName: DataTypes.TEXT,
  favoriteColor: {
    type: DataTypes.TEXT,
    defaultValue: 'green',
  },
  age: DataTypes.INTEGER,
  cash: DataTypes.INTEGER,
});

(async () => {
  await sequelize.sync({ force: true });
  // Code here
})();

export default User;