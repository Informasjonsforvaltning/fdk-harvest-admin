import { Router } from 'express';
import dataSourceHandlers from './data-source.handlers';
import { dataSourceApiValidator } from './data-source.validator';
import { enforceAuthority } from './keycloak';

export const dataSourcesPath = '/datasources';

export const createDataSourceRouter = (): Router => {
  const dataSourceRouter = Router();

  dataSourceRouter.get(
    `${dataSourcesPath}`,
    dataSourceApiValidator.validate('get', `${dataSourcesPath}`),
    dataSourceHandlers.getAll
  );

  dataSourceRouter.post(
    `${dataSourcesPath}`,
    enforceAuthority(),
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
    enforceAuthority(),
    dataSourceApiValidator.validate('put', `${dataSourcesPath}/{id}`),
    dataSourceHandlers.update
  );

  dataSourceRouter.delete(
    `${dataSourcesPath}/:id`,
    enforceAuthority(),
    dataSourceApiValidator.validate('delete', `${dataSourcesPath}/{id}`),
    dataSourceHandlers.deleteById
  );

  dataSourceRouter.post(
    `${dataSourcesPath}/:id`,
    enforceAuthority(),
    dataSourceApiValidator.validate('post', `${dataSourcesPath}/{id}`),
    dataSourceHandlers.harvestById
  );

  return dataSourceRouter;
};
