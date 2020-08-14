import yjs from 'js-yaml';
import fs from 'fs';

export const readSyncYaml = (path: string): any =>
  yjs.safeLoad(fs.readFileSync(path, 'utf-8'));
