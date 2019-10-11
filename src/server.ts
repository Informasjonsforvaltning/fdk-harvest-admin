import config = require('config');

import { createApp } from './app';
import { createMessageBroker, MessageBroker } from './rabbitmq/rabbitmq';

const PORT = config.get('server.port') || 8000;

const { host, port, name, username = '', password = '' } = config.get(
  'mongodb'
);

const connectionUris = `mongodb://${username}:${password}@${host}:${port}/${name}?authSource=admin&authMechanism=SCRAM-SHA-1`;

createMessageBroker()
  .then((messageBroker: MessageBroker) =>
    createApp({ connectionUris, messageBroker })
  )
  .then(app => {
    app.listen(Number(PORT), err => {
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
