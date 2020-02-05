// In this file you can configure migrate-mongo

const c = require('config');

const { host, port, name, username = '', password = '' } = c.get('mongodb');

const config = {
  mongodb: {
    // TODO Change (or review) the url to your MongoDB:
    url: `mongodb://${username}:${password}@${host}:${port}/${name}?authSource=admin&authMechanism=SCRAM-SHA-1`,

    // TODO Change this to your database name:
    databaseName: name,

    options: {
      useNewUrlParser: true // removes a deprecation warning when connecting
      // useUnifiedTopology: true // removes a deprecating warning when connecting
      //   connectTimeoutMS: 3600000, // increase connection timeout to 1 hour
      //   socketTimeoutMS: 3600000, // increase socket timeout to 1 hour
    }
  },

  // The migrations dir, can be an relative or absolute path. Only edit this when really necessary.
  migrationsDir: 'migrations',

  // The mongodb collection where the applied changes are stored. Only edit this when really necessary.
  changelogCollectionName: 'changelog'
};

// Return the config as a promise
module.exports = config;
