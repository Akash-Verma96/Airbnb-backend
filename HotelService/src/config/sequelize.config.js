require('ts-node/register'); // Dynamically register TypeScript for Node.js now it will convert in js files on the fly
const config = require('./db.config');
module.exports = config;
