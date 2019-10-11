/* eslint-disable @typescript-eslint/no-explicit-any */
import amqplib, { Channel, Connection } from 'amqplib';
import config from 'config';
import { DataSourceDocument } from '../data-source.model';

interface HarvestCatalogueMessage {
  publisherId: string;
}

export interface MessageBroker {
  publishDataSource: (doc: DataSourceDocument) => void;
}

export const createMessageBroker = async (): Promise<MessageBroker> => {
  const { user, pass, host, port, exchange } = config.get('rabbitmq');
  const connectionUri = `amqp://${user}:${pass}@${host}:${port}`;

  const connection: Connection = await amqplib.connect(connectionUri);
  connection.on('error', console.error);
  const channel: Channel = await connection.createChannel();
  channel.assertExchange(exchange, 'topic', { durable: false });

  return {
    publishDataSource: ({ publisherId = '' }: DataSourceDocument): void => {
      const message: HarvestCatalogueMessage = {
        publisherId: publisherId
      };

      channel.publish(
        exchange,
        'conceptPublisher.HarvestTrigger',
        Buffer.from(JSON.stringify(message))
      );
    }
  };
};
