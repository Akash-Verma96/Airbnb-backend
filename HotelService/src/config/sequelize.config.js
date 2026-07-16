require('ts-node/register'); // Dynamically register TypeScript for Node.js now it will convert in js files on the fly
const config = require('./db.config');
module.exports = config;


// require('ts-node/register'); // This line enables TypeScript support
// const config = require('./db.config');
// module.exports = config;