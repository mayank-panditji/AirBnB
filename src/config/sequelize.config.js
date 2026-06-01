require('ts-node/register');//this line enables typescript support beacuse sequelize not support the typescript directly
const config = require('./db.config').default;
module.exports = config;