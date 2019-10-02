/* eslint-disable @typescript-eslint/no-explicit-any */
import express, { Request, Response, NextFunction } from 'express';
import dotenv from 'dotenv';
import path from 'path';
import { path as approot } from 'app-root-path';
import swaggerUI from 'swagger-ui-express';
import { Utils } from './utils/utils';

dotenv.config({ path: path.join(approot, '.env') });

import { connect } from './db/db';
import { DatasourceRouter } from './routes/routers';

const { PORT = 8000 } = process.env;

const app = express();

const commonErrorHandler = (
  err: any,
  _req: Request,
  res: Response,
  _next: NextFunction
): void => {
  res.status(err.statusCode).json({
    message: err.message,
    errors: err.errors
  });
};

app.use(express.json());
app.use(commonErrorHandler);

app.use(
  '/api-docs',
  swaggerUI.serve,
  swaggerUI.setup(Utils.readOpenApi(false))
);

app.use('/api/datasources', DatasourceRouter);

if (process.env.NODE_ENV === 'test') {
  app.listen(PORT as number, Utils.listenCallback);
} else {
  connect()
    .then(() => {
      app.listen(PORT as number, Utils.listenCallback);
    })
    .catch(Utils.abortAndExit);
}

export default app;
