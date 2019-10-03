import { Router } from 'express';
import { dataSourceHandlers } from './data-source.handlers';

export const dataSourcesPath = '/datasources';
export const dataSourceRouter = Router();

dataSourceRouter.get(`${dataSourcesPath}`, dataSourceHandlers.getAll);

dataSourceRouter.post(`${dataSourcesPath}`, dataSourceHandlers.create);

dataSourceRouter.get(`${dataSourcesPath}/:id`, dataSourceHandlers.getById);

dataSourceRouter.put(`${dataSourcesPath}/:id`, dataSourceHandlers.update);
