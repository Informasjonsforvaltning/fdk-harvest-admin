import AsyncApiValidator from 'asyncapi-validator';
import { resolve } from 'app-root-path';
import config from 'config';

import { readSyncYaml } from '../lib/read-sync-yaml';

const asyncSpecFile = resolve(`/spec/${config.get('spec.async-api')}`);

export const validator = AsyncApiValidator.fromSource(
  readSyncYaml(asyncSpecFile)
);
