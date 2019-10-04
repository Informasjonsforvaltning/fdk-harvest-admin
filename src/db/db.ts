import mongoose from 'mongoose';

export const connectDb = async (connectionUris: string): Promise<void> => {
  const options = {
    // fix deprecations
    useUnifiedTopology: true,
    useNewUrlParser: true,
    useFindAndModify: false,
    useCreateIndex: true
  };

  await mongoose.connect(connectionUris, options);
  mongoose.connection.on('error', console.error);
};
