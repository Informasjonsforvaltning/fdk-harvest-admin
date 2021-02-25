import express, { Application } from 'express';
import morgan from 'morgan';
import cors from 'cors';
import bodyParser from 'body-parser';

import { commonErrorHandler } from './lib/common-error-handler';
import { connectDb } from './db/db';
import { createDataSourceRouter } from './data-source.router';
import keycloak from './keycloak';
import { streamLogger } from './logger';

export async function createApp({
  connectionUris
}: {
  connectionUris: string;
}): Promise<Application> {
  const app = express();

  app.use(bodyParser.json());

  app.use(cors({ origin: '*', exposedHeaders: 'Location' }));

  app.use(keycloak.middleware());

  app.use('/api', createDataSourceRouter());

  app.use(commonErrorHandler);

  app.use(
    morgan('combined', {
      skip: function (req: any) {
        return req.url === '/ping' || req.url === '/ready';
      },
      stream: streamLogger
    })
  );

  await connectDb(connectionUris);
  return app;
}
