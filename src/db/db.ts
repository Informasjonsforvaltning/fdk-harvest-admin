import mongoose from 'mongoose';
import config from 'config';

// export const close = async (): Promise<void> => {
//      await mongoose.disconnect()
// };

export const connect = async (): Promise<void> => {
  const { host, port, name } = config.get('mongodb');

  const url = `mongodb://${process.env.MONGO_USERNAME || ''}:${process.env
    .MONGO_PASSWORD || ''}@${host}:${port}/${name}`;

  const options = {
    // fix deprecations
    useUnifiedTopology: true,
    useNewUrlParser: true,
    useFindAndModify: false,
    useCreateIndex: true
  };

  await mongoose.connect(url, options);
  console.log('main db connected', url)
  mongoose.connection.on('error', console.error);
};
