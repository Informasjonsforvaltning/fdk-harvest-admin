import express, { Application } from 'express';
import bodyParser from 'body-parser';
import { commonErrorHandler } from './lib/common-error-handler';

import { connect } from './db/db';
import { dataSourceRouter } from './data-source.router';

export async function createApp(): Promise<Application> {
  const app = express();

  app.use(bodyParser.json());

  app.use('/api', dataSourceRouter);

  app.use(commonErrorHandler);

  await connect();
  return app;
}
