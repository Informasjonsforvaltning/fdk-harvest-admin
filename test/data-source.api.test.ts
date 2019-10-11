/* eslint-disable @typescript-eslint/no-explicit-any */
import request from 'supertest';
import { createApp } from '../src/app';

import { MongoMemoryServer } from 'mongodb-memory-server';
import mongoose from 'mongoose';

import { dataSourceApiValidator, spec } from '../src/data-source.validator';
import { Application } from 'express';
import { MessageBroker } from '../src/rabbitmq/rabbitmq';

const messageBrokerMock: MessageBroker = {
  publishDataSource: (): void => {}
};

const mongoTestServer = new MongoMemoryServer();

let app: Application;

beforeEach(async () => {
  const connectionUris = await mongoTestServer.getConnectionString();
  const messageBroker: MessageBroker = messageBrokerMock;
  app = await createApp({ connectionUris, messageBroker });
});

afterEach(async () => mongoose.disconnect());

describe('/api/datasources', () => {
  const supportedMethods = spec.paths['/datasources'];

  const { post = {} as any } = supportedMethods;
  it(post.description || post.summary, async () => {
    const { responses } = post;
    const code = Object.keys(responses)[0];
    const uuidV4regExp = new RegExp(
      /^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i
    );

    await request(app)
      .post(`/api/datasources`)
      .send({
        dataSourceType: 'SKOS-AP-NO',
        url: 'http://example.com',
        acceptHeaderValue: 'text/turtle',
        publisherId: 'hi',
        description: 'Bla bla bla'
      })
      .expect(parseInt(code))
      .expect('Location', uuidV4regExp)
      .expect(dataSourceApiValidator.validateResponse('post', '/datasources'));
  });
});
