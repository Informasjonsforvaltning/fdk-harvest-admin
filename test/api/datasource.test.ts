import request from 'supertest';
import app from '../../src/app';

import { MongoMemoryServer } from 'mongodb-memory-server';
import mongoose from 'mongoose';

import {spec, dataSourceApiValidator} from '../../src/data-source.validator';

beforeEach(done => {
  mongoose.Promise = global.Promise;
  const mongoTestServer = new MongoMemoryServer();

  const options = {
    autoReconnect: true,
    reconnectTries: Number.MAX_VALUE,
    reconnectInterval: 1000,
    useUnifiedTopology: true,
    useNewUrlParser: true,
    useFindAndModify: false,
    useCreateIndex: true
  };

  mongoTestServer
    .getConnectionString()
    .then(uri => {
      mongoose.connection.once('open', () => {
        console.log(`MongoDB successfully connected to ${uri}`);
      });
      mongoose
        .connect(uri, options)
        .then(() => {
          mongoose.connection.on('error', e => {
            if (e.message.code === 'ETIMEDOUT') {
              console.log(e);
              mongoose.connect(uri, options);
            }
            console.log(e);
          });
          done();
        })
        .catch(done);
    })
    .catch(done);
});

afterEach(done => {
  mongoose
    .disconnect()
    .then(done)
    .catch(done);
});

describe('/api/datasources', () => {
  const supportedMethods = spec.paths['/datasources'];

  const { post={} as any } = supportedMethods;
  it(post.description || post.summary, async ()=> {
    const { responses } = post;
    const code = Object.keys(responses)[0];
    const uuidV4regExp = new RegExp(/^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i)

    await request(app)
      .post(`/api/datasources`)
      .send({
        id: 'fx434358zx45sdss11',
        dataSourceType: 'SKOS-AP-NO',
        url: 'http://example.com',
        publisherId: 'hi',
        description: 'Bla bla bla'
      })
      .expect(parseInt(code))
      .expect('Location', uuidV4regExp)
      .expect(dataSourceApiValidator.validateResponse('post', '/datasources'))
  });
});
