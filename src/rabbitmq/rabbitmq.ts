/* eslint-disable @typescript-eslint/no-explicit-any */
import amqplib, { Channel, Message } from 'amqplib/callback_api';
import config from 'config';
import { DataSourceDocument, DataSourceModel } from '../data-source.model';
import { validator } from './asyncspec-validator';

interface HarvestCatalogueMessage {
  publisherId: string;
  dataType: string;
}

let channel: Channel | null;

const {
  user,
  pass,
  host,
  port,
  exchange,
  listenerKey,
  publisherPartialKey,
  validationKey
} = config.get('rabbitmq');
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

        ch.assertQueue(
          '',
          {
            exclusive: true
          },
          (err, q) => {
            if (err) {
              console.error('listener queue could not be asserted');
              return connection.close();
            }
            ch.bindQueue(q.queue, exchange, listenerKey);

            ch.consume(q.queue, async ({ content, fields }: Message) => {
              console.log(
                `[*] received new datasource from: ${fields.routingKey}`
              );

              const dataSource = JSON.parse(content.toString());
              try {
                if ((await validator).validate(validationKey, dataSource)) {
                  DataSourceModel.find({
                    publisherId: dataSource.publisherId,
                    url: dataSource.url })
                  .then((docs: DataSourceDocument[]) => {
                    if(!docs.length) {
                      new DataSourceModel(dataSource).save();
                    };
                  });
                }
              } catch (e) {
                console.error(e);
              }
            });

            channel = ch;
          }
        );
      });
    }
  });
};

export const publishDataSource = ({
  id = '',
  dataType = '',
  publisherId = ''
}: DataSourceDocument): void => {
  const message: HarvestCatalogueMessage = {
    publisherId: publisherId,
    dataType: dataType
  };

  const key = `${dataType}.${publisherPartialKey}`;

  if (channel && dataType) {
    console.log(
      `[id:${id}] organization (${publisherId}) publishing to ${key} ...`
    );
    channel.publish(exchange, key, Buffer.from(JSON.stringify(message)), {
      contentType: 'application/json'
    });
  }
};
