import AsyncApiValidator from 'asyncapi-validator';
import { readSyncYaml } from '../lib/read-sync-yaml';
import config from 'config';

const asyncSpecFile = require('app-root-path').resolve(
  `/spec/${config.get('spec.async-api')}`
);

export const validator = AsyncApiValidator.fromSource(
  readSyncYaml(asyncSpecFile)
);
