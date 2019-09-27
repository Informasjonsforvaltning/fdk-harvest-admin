import mongoose from 'mongoose';
import config from 'config';

mongoose.Promise = global.Promise;

const logRuntimeError = (error: mongoose.Error): void => {
  console.log(error);
};

const abortAndExit = (error: mongoose.Error): void => {
  logRuntimeError(error);
  process.exit(0);
};

export default (): void => {
  const { host, port, name } = config.get('mongodb');
  const {
    FDKHARVESTADM_MONGO_USERNAME = '',
    FDKHARVESTADM_MONGO_PASSWORD = ''
  } = process.env;

  const url = `mongodb://${FDKHARVESTADM_MONGO_USERNAME}:${FDKHARVESTADM_MONGO_PASSWORD}@${host}:${port}/${name}`;

  const options = {
    useNewUrlParser: true,
    useUnifiedTopology: true
  };

  mongoose
    .connect(url, options)
    .then(() => {
      mongoose.connection.on('error', logRuntimeError);
    })
    .catch(abortAndExit);
};
