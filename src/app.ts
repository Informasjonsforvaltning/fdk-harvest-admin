import express from 'express';
import bodyParser from 'body-parser';
import { commonErrorHandler } from './lib/common-error-handler';

import { connect } from './db/db';
import { dataSourceRouter } from './data-source.router';

const app = express();

app.use(bodyParser.json());

app.use('/api', dataSourceRouter);

app.use(commonErrorHandler);

async function main(): Promise<void> {
  await connect();

  app.listen(8000, err => {
    if (err) {
      throw err;
    }
    console.log('server running on: 8000');
  });
}

main().catch(console.error);

export default app;
