/* eslint-disable @typescript-eslint/no-explicit-any */
import yjs from 'js-yaml';
import { path as approot } from 'app-root-path';
import path from 'path';
import fs from 'fs';
import config from 'config';

export class Utils {
  static readOpenApi(fatal = true): any {
    const { filename } = config.get('openapi');
    try {
      return yjs.safeLoad(
        fs.readFileSync(path.join(approot, filename), 'utf-8')
      );
    } catch (e) {
      console.log(e);
      fatal && process.exit(0);
    }
  }

  static listenCallback = (err: any): void => {
    console.log('server listenening ...');
    err && Utils.abortAndExit(err);
  };

  static abortAndExit = (err: any): void => {
    console.log(err);
    process.exit(0);
  };
}
