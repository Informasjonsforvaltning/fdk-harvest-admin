import { createApp } from './app';
import config = require('config');

const { PORT = 8000 } = process.env;

const { host, port, name } = config.get('mongodb');

const connectionUris = `mongodb://${process.env.MONGO_USERNAME || ''}:${process
  .env.MONGO_PASSWORD || ''}@${host}:${port}/${name}`;

createApp({ connectionUris })
  .then(app => {
    app.listen(Number(PORT), err => {
      if (err) {
        throw err;
      }
      console.log(`server running on: ${PORT}`);
    });
  })
  .catch(console.error);
