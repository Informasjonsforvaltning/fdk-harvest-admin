import mongoose from 'mongoose';
import config from 'config';
import Promise from 'promise';

mongoose.Promise = global.Promise;

const logRuntimeError = (error: mongoose.Error): void => {
  console.log(error);
};

export const close = (): Promise<void> => {
  return new Promise((resolve, reject): void => {
    mongoose
      .disconnect()
      .then(resolve)
      .catch(reject);
  });
};

export const connect = (): Promise<void> => {
  const { host, port, name } = config.get('mongodb');
  const {
    FDKHARVESTADM_MONGO_USERNAME = '',
    FDKHARVESTADM_MONGO_PASSWORD = ''
  } = process.env;

  const url = `mongodb://${FDKHARVESTADM_MONGO_USERNAME}:${FDKHARVESTADM_MONGO_PASSWORD}@${host}:${port}/${name}`;

  const options = {
    // fix deprecations
    useUnifiedTopology: true,
    useNewUrlParser: true,
    useFindAndModify: false,
    useCreateIndex: true
  };

  return new Promise((resolve, reject): void => {
    mongoose
      .connect(url, options)
      .then(() => {
        mongoose.connection.on('error', logRuntimeError);
        resolve();
      })
      .catch(err => {
        logRuntimeError(err);
        reject(err);
      });
  });
};
