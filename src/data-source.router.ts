import { Router } from 'express';
import { dataSourceHandlers } from './data-source.handlers';
import { dataSourceApiValidator } from './data-source.validator';

export const dataSourcesPath = '/datasources';
export const dataSourceRouter = Router();

dataSourceRouter.get(
  `${dataSourcesPath}`,
  dataSourceApiValidator.validate('get', `${dataSourcesPath}`),
  dataSourceHandlers.getAll
);

dataSourceRouter.post(
  `${dataSourcesPath}`,
  dataSourceApiValidator.validate('post', `${dataSourcesPath}`),
  dataSourceHandlers.create
);

dataSourceRouter.get(
  `${dataSourcesPath}/:id`,
  dataSourceApiValidator.validate('get', `${dataSourcesPath}/{id}`),
  dataSourceHandlers.getById
);

dataSourceRouter.put(
  `${dataSourcesPath}/:id`,
  dataSourceApiValidator.validate('put', `${dataSourcesPath}/{id}`),
  dataSourceHandlers.update
);
