import { OpenApiDocument, OpenApiValidator } from 'express-openapi-validate';
import { readSyncYaml } from './lib/read-sync-yaml';
import config from 'config';

const specPath = require('app-root-path').resolve(
  `/spec/${config.get('spec.open-api')}`
);
export const spec: OpenApiDocument = readSyncYaml(specPath) as OpenApiDocument;

export const dataSourceApiValidator = new OpenApiValidator(spec);
