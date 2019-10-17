import express, { Application } from 'express';
import cors from 'cors';
import bodyParser from 'body-parser';

import { commonErrorHandler } from './lib/common-error-handler';
import { connectDb } from './db/db';
import { createDataSourceRouter } from './data-source.router';
import { MessageBroker } from './rabbitmq/rabbitmq';
import config from 'config';
import Keycloak from 'keycloak-connect';

export async function createApp({
  connectionUris,
  messageBroker
}: {
  connectionUris: string;
  messageBroker: MessageBroker;
}): Promise<Application> {
  const app = express();
  const keycloak = new Keycloak({}, config.get('keycloak'));

  app.use(bodyParser.json());

  app.use(cors({ origin: '*', exposedHeaders: 'Location' }));

  app.use(keycloak.middleware());

  app.use('/api', createDataSourceRouter(messageBroker, keycloak));

  app.use(commonErrorHandler);

  await connectDb(connectionUris);
  return app;
}
