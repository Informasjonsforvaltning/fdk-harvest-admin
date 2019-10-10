/* eslint-disable @typescript-eslint/no-explicit-any */
import amqplib, { Channel } from 'amqplib/callback_api';
import config from 'config';
import { DataSourceDocument } from '../data-source.model';

interface HarvestCatalogueMessage {
  publisherId: string;
  catalogueId: string;
  dataSourceType: string;
}

const { user, pass, host, port, exchange, topic } = config.get('rabbitmq');

let channel: Channel;

amqplib.connect(`amqp://${user}:${pass}@${host}:${port}`, (err, connection) => {
  if (err) {
    throw err;
  }

  connection.on('error', console.error);

  connection.createChannel((err, ch) => {
    if (err) {
      throw err;
    }

    ch.assertExchange(exchange, 'direct', { durable: false });
    channel = ch;
  });
});

export const publishDataSource = ({
  dataSourceType = '',
  publisherId = ''
}: DataSourceDocument): void => {
  const message: HarvestCatalogueMessage = {
    publisherId: publisherId,
    catalogueId: publisherId,
    dataSourceType: dataSourceType
  };

  channel.publish(exchange, topic, Buffer.from(JSON.stringify(message)));
};

process.on('exit', () => {
  channel.close(console.log);
});
