import request from 'supertest';
import { expect } from 'chai';
import app from '../../src/app';
import yjs from 'js-yaml';
import fs from 'fs';
import { path as approot } from 'app-root-path';
import path from 'path';

const CONTENT_TYPE_JSON = 'application/json';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
let spec: any;
try {
  spec = yjs.safeLoad(
    fs.readFileSync(path.join(approot, 'fdk-harvest-adm.yaml'), 'utf-8')
  );
} catch (e) {
  console.log(e);
}

// TODO(chlenix): replace with AJV schema validator
// eslint-disable-next-line @typescript-eslint/no-explicit-any
const validateDatasourceSchema = (datasource: any): void => {
  const schema = spec.components.schemas.DataSource;
  expect(datasource).to.be.an(schema.type);
  expect(datasource).to.have.all.keys(...Object.keys(schema.properties));
};

describe('/api/datasources', () => {
  const supportedMethods = spec.paths['/datasources'];
  const requestMeta = supportedMethods.get;

  it(requestMeta.description, () => {
    const responses = requestMeta.responses;
    const code = Object.keys(responses)[0];
    request(app)
      .get(`/api/datasources`)
      .expect(parseInt(code, 10))
      .then(res => {
        expect(res.body).to.be.an(
          responses[code].content[CONTENT_TYPE_JSON].schema.type
        );

        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        res.body.map(validateDatasourceSchema);
      })
      .catch(err => expect(err).to.be.undefined);
  });
});
