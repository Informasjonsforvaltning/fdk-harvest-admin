import { OpenApiDocument, OpenApiValidator } from 'express-openapi-validate';
import { resolve } from 'app-root-path';
import config from 'config';

import { readSyncYaml } from './lib/read-sync-yaml';

const specPath = resolve(`/spec/${config.get('spec.open-api')}`);

export const spec: OpenApiDocument = readSyncYaml(specPath) as OpenApiDocument;

export const dataSourceApiValidator = new OpenApiValidator(spec);
