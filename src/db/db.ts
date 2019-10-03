import mongoose from 'mongoose';
import config from 'config';

export const connect = async (): Promise<void> => {
  const { host, port, name } = config.get('mongodb');

  const url = `mongodb://${process.env.MONGO_USERNAME || ''}:${process.env
    .MONGO_PASSWORD || ''}@${host}:${port}/${name}`;

  const options = {
    useNewUrlParser: true,
    useUnifiedTopology: true
  };

  await mongoose.connect(url, options);
  mongoose.connection.on('error', console.error);
};
