import express, { Application } from 'express';
import cors from 'cors';
import bodyParser from 'body-parser';
import { commonErrorHandler } from './lib/common-error-handler';

import { connectDb } from './db/db';
import { createDataSourceRouter } from './data-source.router';

export async function createApp({
  connectionUris
}: {
  connectionUris: string;
}): Promise<Application> {
  const app = express();

  app.use(bodyParser.json());

  app.use(cors({ origin: '*', exposedHeaders: 'Location' }));

  app.use('/api', createDataSourceRouter());

  app.use(commonErrorHandler);

  await connectDb(connectionUris);
  return app;
}
