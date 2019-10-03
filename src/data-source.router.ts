import { Router } from 'express';
import { dataSourceHandlers } from './data-source.handlers';

export const dataSourcesPath = '/api/datasources';
export const dataSourceRouter = Router();

dataSourceRouter.get(`/`, dataSourceHandlers.getAll);
dataSourceRouter.post(`/`, dataSourceHandlers.create);

dataSourceRouter.get(`/:id`, dataSourceHandlers.getById);
