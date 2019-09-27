import express from 'express';
import bodyParser from 'body-parser';
import dotenv from 'dotenv';
import path from 'path';
import { path as approot } from 'app-root-path';

dotenv.config({ path: path.join(approot, '.env') });

import dbInit from './db/db';
dbInit();

import generateDocs from './docs/swagger-docs';
import { DatasourceRouter } from './routes/routers';

const app = express();

generateDocs(app);

app.use(bodyParser.json());
app.use('/api/datasources', DatasourceRouter);

app.use((
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  err: any,
  _req: express.Request,
  res: express.Response,
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  _next: express.NextFunction
) => {
  // format error
  res.status(err.status).json({
    message: err.message,
    errors: err.errors
  });
});

app.listen(8000, err => {
  console.log('server running');
  if (err) {
    return console.log(err);
  }
});

export default app;
