/* eslint-disable @typescript-eslint/no-explicit-any */
import request from 'supertest';
import { createApp } from '../src/app';

import { MongoMemoryServer } from 'mongodb-memory-server';
import mongoose from 'mongoose';

import { dataSourceApiValidator, spec } from '../src/data-source.validator';
import { Application } from 'express';
import { MessageBroker } from '../src/rabbitmq/rabbitmq';

import { internet, random } from 'faker';

const messageBrokerMock: MessageBroker = {
  publishDataSource: (): void => {}
};

const mongoTestServer = new MongoMemoryServer();

const dataSourceMock = {
  dataSourceType: random.boolean ? 'SKOS-AP-NO' : 'DCAT_AP_NO',
  url: internet.url(),
  acceptHeaderValue: 'text/turtle',
  publisherId: random.number({ min: 800_000_000, max: 900_000_000 }).toString(),
  description: 'descriptive text here'
};

let app: Application;

beforeEach(async () => {
  const messageBroker: MessageBroker = messageBrokerMock;
  const connectionUris = await mongoTestServer.getConnectionString();
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
      .send(dataSourceMock)
      .expect(parseInt(code))
      .expect('Location', uuidV4regExp)
      .expect(dataSourceApiValidator.validateResponse('post', '/datasources'));
  });
});
