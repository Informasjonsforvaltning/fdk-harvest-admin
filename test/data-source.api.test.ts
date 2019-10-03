/* eslint-disable @typescript-eslint/no-explicit-any */
import request from 'supertest';
import { createApp } from '../src/app';

import { MongoMemoryServer } from 'mongodb-memory-server';
import mongoose from 'mongoose';

import { spec, dataSourceApiValidator } from '../src/data-source.validator';
import { Application } from 'express';

const mongoTestServer = new MongoMemoryServer();

let appPromise: Promise<Application>;

beforeEach(() => {
  appPromise = mongoTestServer
    .getConnectionString()
    .then(connectionUris => createApp({ connectionUris }));
});

afterEach(done => {
  mongoose
    .disconnect()
    .then(done)
    .catch(done);
});

describe('/api/datasources', () => {
  const supportedMethods = spec.paths['/datasources'];

  const { post = {} as any } = supportedMethods;
  it(post.description || post.summary, async () => {
    const { responses } = post;
    const code = Object.keys(responses)[0];
    const uuidV4regExp = new RegExp(
      /^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i
    );

    await request(await appPromise)
      .post(`/api/datasources`)
      .send({
        dataSourceType: 'SKOS-AP-NO',
        url: 'http://example.com',
        publisherId: 'hi',
        description: 'Bla bla bla'
      })
      .expect(parseInt(code))
      .expect('Location', uuidV4regExp)
      .expect(dataSourceApiValidator.validateResponse('post', '/datasources'));
  });
});
