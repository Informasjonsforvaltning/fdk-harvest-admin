import config = require('config');

import { createApp } from './app';
import {
  createMessageBroker,
  MessageBroker,
  MOCK_MESSAGE_BROKER
} from './rabbitmq/rabbitmq';

const PORT = config.get('server.port') || 8000;

const { host, port, name, username = '', password = '' } = config.get(
  'mongodb'
);

const { user, pass, rabbitHost, rabbitPort } = config.get('rabbitmq');

const connectionUris = `mongodb://${username}:${password}@${host}:${port}/${name}?authSource=admin&authMechanism=SCRAM-SHA-1`;
const rabbitConnectionUri = `amqp://${user}:${pass}@${rabbitHost}:${rabbitPort}`;

createMessageBroker(rabbitConnectionUri)
  .then((messageBroker: MessageBroker = MOCK_MESSAGE_BROKER) => {
    createApp({ connectionUris, messageBroker }).then(app => {
      app.listen(Number(PORT), err => {
        if (err) {
          throw err;
        }
        console.log(`server running on: ${PORT}`);
      });
    });
  })
  .catch(console.error);
