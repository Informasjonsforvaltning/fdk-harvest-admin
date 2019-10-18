import { Router } from 'express';
import { createDataSourceHandlers } from './data-source.handlers';
import { dataSourceApiValidator } from './data-source.validator';
import { MessageBroker } from './rabbitmq/rabbitmq';
import keycloak from './keycloak';

export const dataSourcesPath = '/datasources';

export const createDataSourceRouter = (
  messageBroker: MessageBroker
): Router => {
  const dataSourceRouter = Router();
  const dataSourceHandlers = createDataSourceHandlers(messageBroker);

  dataSourceRouter.get(
    `${dataSourcesPath}`,
    dataSourceApiValidator.validate('get', `${dataSourcesPath}`),
    dataSourceHandlers.getAll
  );

  dataSourceRouter.post(
    `${dataSourcesPath}`,
    keycloak.protect(),
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
    keycloak.protect(),
    dataSourceApiValidator.validate('put', `${dataSourcesPath}/{id}`),
    dataSourceHandlers.update
  );

  dataSourceRouter.delete(
    `${dataSourcesPath}/:id`,
    keycloak.protect(),
    dataSourceApiValidator.validate('delete', `${dataSourcesPath}/{id}`),
    dataSourceHandlers.deleteById
  );

  dataSourceRouter.post(
    `${dataSourcesPath}/:id`,
    keycloak.protect(),
    dataSourceApiValidator.validate('post', `${dataSourcesPath}/{id}`),
    dataSourceHandlers.harvestById
  );

  return dataSourceRouter;
};
