import mongoose from 'mongoose';
import logger from '../logger';

export const connectDb = async (connectionUris: string): Promise<void> => {
  await mongoose.connect(connectionUris, {});
  mongoose.connection.on('error', logger.error);
};
