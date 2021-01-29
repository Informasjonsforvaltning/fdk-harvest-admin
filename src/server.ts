import config = require('config');

import { createApp } from './app';
import { rabbitConnect } from './rabbitmq/rabbitmq';

const PORT = config.get('server.port') || 8000;

const { host, port, name, username = '', password = '' } = config.get(
  'mongodb'
);

const connectionUris = `mongodb://${username}:${password}@${host}:${port}/${name}?authSource=admin&authMechanism=SCRAM-SHA-1`;

rabbitConnect();

createApp({ connectionUris })
  .then(app => {
    app.listen(Number(PORT)).on('error', err => {
      if (err) {
        throw err;
      }
      console.log(`server running on: ${PORT}`);
    });
  })
  .catch(err => {
    console.error(err);
    process.exit(1);
  });
