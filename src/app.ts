import express from 'express';
import bodyParser from 'body-parser';
import { commonErrorHandler } from './lib/common-error-handler';

const app = express();

app.use(bodyParser.json());
app.use(commonErrorHandler);

app.listen(8000, err => {
  console.log('server running');
  if (err) {
    return console.log(err);
  }
});

export default app;
