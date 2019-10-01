import request from 'supertest';
import { expect } from 'chai';
import app from '../../src/app';

import { Utils } from '../../src/utils/utils';
import { OpenApiValidator } from 'express-openapi-validate';
import { MongoMemoryServer } from 'mongodb-memory-server';
import mongoose from 'mongoose';

const openApiDocument = Utils.readOpenApi();
const validator = new OpenApiValidator(openApiDocument);

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
  const supportedMethods = openApiDocument.paths['/datasources'];

  const { post } = supportedMethods;
  it(post.description || post.summary, done => {
    const { responses } = post;
    const code = Object.keys(responses)[0];
    const validateResponse = validator.validateResponse('post', '/datasources');

    const url = 'http://example.com';

    request(app)
      .post(`/api/datasources`)
      .send({
        id: 'fx434358zx45sdss11',
        dataSourceType: 'SKOS-AP-NO',
        url: url,
        publisher: 'hi',
        description: 'Bla bla bla'
      })
      .expect(parseInt(code))
      .expect('Location', url)
      .then(res => {
        expect(validateResponse(res));
        done();
      })
      .catch(err => expect(err).to.be.undefined);
  });
});
