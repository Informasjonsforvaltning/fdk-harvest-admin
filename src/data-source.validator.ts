import { OpenApiDocument, OpenApiValidator } from 'express-openapi-validate';
import { readSyncYaml } from './lib/read-sync-yaml';

const specPath = __dirname + '/../spec/fdk-harvest-admin.yaml';
export const spec: OpenApiDocument = readSyncYaml(specPath) as OpenApiDocument;

export const dataSourceApiValidator = new OpenApiValidator(spec);
