import express from 'express';
import bodyParser from 'body-parser';
import { commonErrorHandler } from './lib/common-error-handler';

import { connect } from './db/db';
import { dataSourceRouter } from './data-source.router';
import config from "config";

const { PORT = 8000 } = process.env;

const app = express();

app.use(bodyParser.json());

app.use('/api', dataSourceRouter);

app.use(commonErrorHandler);

async function main(): Promise<void> {
  const { listen } = config.get('server');
  const database = config.get('database');
  if (database!=="inMemory") {await connect()};

  if(listen) {
    app.listen(Number(PORT), err => {
      if (err) {
        throw err;
      }
      console.log(`server running on: ${PORT}`);
    });
  }
}

main().catch(console.error);

export default app;
