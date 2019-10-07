/* eslint-disable @typescript-eslint/no-explicit-any */
import amqplib, { Channel } from 'amqplib/callback_api';
import config from 'config';

export interface HarvestMessage {
  orgId: string;
  catalogId: string;
  datasourceType: string;
}

const { host = '', port = '' } = config.get('rabbitmq');
let ch: Channel;

const bail = (err: any): void => {
  console.error(err);
  process.exit(0);
};

amqplib.connect(`amqp://${host}:${port}`, (err, connection) => {
  err && bail(err);

  connection.on('error', console.error);

  connection.createChannel((err, channel) => {
    err && bail(err);
    ch = channel;
  });
});

export const publishToQueue = async (
  payload: HarvestMessage
): Promise<void> => {
  const envelope = {
    type: `${payload.datasourceType.toUpperCase()}_HARVEST_TRIGGER`,
    payload: payload
  };
  ch.sendToQueue('test-q', Buffer.from(JSON.stringify(envelope), 'utf-8'));
};

process.on('exit', () => {
  ch.close(console.log);
});
