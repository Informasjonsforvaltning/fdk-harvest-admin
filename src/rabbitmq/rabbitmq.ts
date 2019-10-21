/* eslint-disable @typescript-eslint/no-explicit-any */
import amqplib, { Channel } from 'amqplib/callback_api';
import config from 'config';
import { DataSourceDocument } from '../data-source.model';

interface HarvestCatalogueMessage {
  publisherId: string;
}

let channel: Channel | null;

const { user, pass, host, port, exchange } = config.get('rabbitmq');
const connectionUri = `amqp://${user}:${pass}@${host}:${port}`;

export const rabbitConnect = (): void => {
  amqplib.connect(connectionUri, (e, connection) => {
    if (e) {
      if (e.code === 'ECONNREFUSED') {
        console.error(
          `${e.code}: unable to connect to rabbit at ${e.address}:${e.port}`
        );
      }
      setTimeout(rabbitConnect, 5000);
    } else {
      connection.on('error', console.error);
      connection.on('close', () => {
        console.error(`Lost connection to rabbitmq, reconnecting ...`);
        channel = null;
        rabbitConnect();
      });
      connection.createChannel((err, ch) => {
        if (err) {
          return connection.close();
        }
        ch.assertExchange(exchange, 'topic', { durable: false });
        channel = ch;
      });
    }
  });
};

export const publishDataSource = ({
  publisherId = ''
}: DataSourceDocument): void => {
  const message: HarvestCatalogueMessage = {
    publisherId: publisherId
  };

  channel &&
    channel.publish(
      exchange,
      'conceptPublisher.HarvestTrigger',
      Buffer.from(JSON.stringify(message)),
      { contentType: 'application/json' }
    );
};
