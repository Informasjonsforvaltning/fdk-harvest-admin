/* eslint-disable @typescript-eslint/no-explicit-any */
import request from 'supertest';
import mongoose from 'mongoose';
import { Application } from 'express';
import { MongoMemoryServer } from 'mongodb-memory-server';
import { internet, random } from 'faker';
import { stub } from 'sinon';
import { NextFunction } from 'connect';

import keycloak from '../src/keycloak';
import { createApp } from '../src/app';
import { MessageBroker } from '../src/rabbitmq/rabbitmq';
import { DataSourceModel } from '../src/data-source.model';
import { dataSourceApiValidator, spec } from '../src/data-source.validator';

const messageBrokerMock: MessageBroker = {
  publishDataSource: (): void => {}
};

const mongoTestServer = new MongoMemoryServer();

const uuidV4regExp = new RegExp(
  /^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i
);

interface DataSourceMock {
  id: string;
  dataSourceType: string;
  url: string;
  acceptHeaderValue: string;
  publisherId: string;
  description: string;
}

const generateDataSourceMock = (): DataSourceMock => {
  return {
    id: random.uuid(),
    dataSourceType: random.boolean ? 'SKOS-AP-NO' : 'DCAT_AP_NO',
    url: internet.url(),
    acceptHeaderValue: 'text/turtle',
    publisherId: random
      .number({ min: 800_000_000, max: 999_999_999 })
      .toString(),
    description: 'descriptive text here'
  };
};

const middlewareMock = (
  _req: Request,
  _res: Response,
  next: NextFunction
): void => {
  next();
};

let app: Application;

before(async () => {
  const messageBroker: MessageBroker = messageBrokerMock;
  const connectionUris = await mongoTestServer.getConnectionString();
  stub(keycloak, 'protect').callsFake((): any => middlewareMock);

  app = await createApp({ connectionUris, messageBroker });
});

afterEach(async () => {
  await DataSourceModel.deleteMany({});
});

after(async () => {
  await mongoose.disconnect();
});

/* POST /api/datasources */
describe('/api/datasources', () => {
  const supportedMethods = spec.paths['/datasources'];

  const { post = {} as any, get = {} as any } = supportedMethods;
  it(post.description || post.summary, async () => {
    const { responses } = post;
    const code = Object.keys(responses)[0];

    await request(app)
      .post(`/api/datasources`)
      .send(generateDataSourceMock())
      .expect(parseInt(code))
      .expect('Location', uuidV4regExp)
      .expect(dataSourceApiValidator.validateResponse('post', '/datasources'));
  });

  it(get.description || get.summary, async () => {
    const { responses } = get;
    const code = Object.keys(responses)[0];

    await request(app)
      .get(`/api/datasources`)
      .expect(parseInt(code))
      .expect(dataSourceApiValidator.validateResponse('get', '/datasources'));
  });
});

describe('/api/datasources/{id}', () => {
  const supportedMethods = spec.paths['/datasources/{id}'];

  let dataSourceMock: any = {};
  beforeEach(async () => {
    dataSourceMock = generateDataSourceMock();
    await new DataSourceModel(dataSourceMock).save();
  });

  const {
    get = {} as any,
    put = {} as any,
    post = {} as any,
    delete: del = {} as any
  } = supportedMethods;

  it(del.description || del.summary, async () => {
    const { responses } = del;
    const code = Object.keys(responses)[0];

    await request(app)
      .delete(`/api/datasources/${dataSourceMock.id}`)
      .expect(parseInt(code))
      .expect(
        dataSourceApiValidator.validateResponse('delete', '/datasources/{id}')
      );
  });

  it(put.description || put.summary, async () => {
    const { responses } = put;
    const code = Object.keys(responses)[0];

    await request(app)
      .put(`/api/datasources/${dataSourceMock.id}`)
      .send(dataSourceMock)
      .expect(parseInt(code))
      .expect(
        dataSourceApiValidator.validateResponse('put', '/datasources/{id}')
      );
  });

  it(get.description || get.summary, async () => {
    const { responses } = get;
    const code = Object.keys(responses)[0];

    await request(app)
      .get(`/api/datasources/${dataSourceMock.id}`)
      .expect(parseInt(code))
      .expect(
        dataSourceApiValidator.validateResponse('get', '/datasources/{id}')
      );
  });

  it(post.description || post.summary, async () => {
    const { responses } = post;
    const code = Object.keys(responses)[0];

    await request(app)
      .post(`/api/datasources/${dataSourceMock.id}`)
      .expect(parseInt(code))
      .expect(
        dataSourceApiValidator.validateResponse('post', '/datasources/{id}')
      );
  });
});
