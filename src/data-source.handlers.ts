import { RequestHandler } from 'express';
import omit from 'lodash/omit';

import { DataSourceDocument, DataSourceModel } from './data-source.model';
import { NotFoundHttpError } from './lib/http-error';
import { elseThrow } from './lib/else-throw';
import { publishDataSource } from './rabbitmq/rabbitmq';
import logger from './logger';

interface ResourceHandlerMap {
  create: RequestHandler;
  getById: RequestHandler;
  getAll: RequestHandler;
  update: RequestHandler;
  deleteById: RequestHandler;
  harvestById: RequestHandler;
}

export default {
  create: (req, res, next): void => {
    const data = omit(req.body, 'id');
    const { allowedOrganizations } = res.locals;

    if (
      allowedOrganizations &&
      allowedOrganizations.length > 0 &&
      !allowedOrganizations.includes(data.publisherId)
    ) {
      res.status(403).send();

      return;
    }

    new DataSourceModel(data)
      .save()
      .then(doc => {
        res.location(doc.id).status(201).send();
        logger.info(`Creating new DataSource with url ${doc.url}`);
        publishDataSource(doc);
      })
      .catch(next);
  },

  update: (req, res, next): void => {
    const { id } = req.params;
    const data = omit(req.body, 'id');
    const { allowedOrganizations } = res.locals;

    if (
      allowedOrganizations &&
      allowedOrganizations.length > 0 &&
      !allowedOrganizations.includes(data.publisherId)
    ) {
      res.status(403).send();

      return;
    }

    DataSourceModel.findOneAndUpdate({ id }, data, { new: true })
      .then(elseThrow<DataSourceDocument>(() => new NotFoundHttpError()))
      .then((doc: DataSourceDocument) => {
        res.status(200).send(doc.toObject());
        logger.info(`Updating DataSource with url ${doc.url}`);
        publishDataSource(doc);
      })
      .catch(next);
  },

  getById: (req, res, next): void => {
    const { id } = req.params;
    DataSourceModel.findOne({ id })
      .then(elseThrow<DataSourceDocument>(() => new NotFoundHttpError()))
      .then((doc: DataSourceDocument) => {
        logger.info(`Get DataSource with url ${doc.url}`);
        res.status(200).send(doc.toObject());
      })
      .catch(next);
  },

  getAll: (req, res, next): void => {
    DataSourceModel.find(req.query)
      .then((docs: DataSourceDocument[]) =>
        res.send(docs.map((doc: DataSourceDocument) => doc.toObject()))
      )
      .catch(next);
  },

  deleteById: (req, res, next): void => {
    const { id } = req.params;
    const { allowedOrganizations } = res.locals;

    DataSourceModel.findOne({ id })
      .then(elseThrow<DataSourceDocument>(() => new NotFoundHttpError()))
      .then((doc: DataSourceDocument) => {
        if (
          allowedOrganizations &&
          allowedOrganizations.length > 0 &&
          !allowedOrganizations.includes(doc.toObject().publisherId)
        ) {
          res.status(403).send();

          return;
        }

        DataSourceModel.deleteOne({ id })
          .then(() => {
            logger.info(`Delete DataSource with url ${doc.url}`);
            res.status(204).send();
          })
          .catch(next);
      })
      .catch(next);
  },

  harvestById: (req, res, next): void => {
    const { id } = req.params;
    const { allowedOrganizations } = res.locals;

    DataSourceModel.findOne({ id })
      .then(elseThrow<DataSourceDocument>(() => new NotFoundHttpError()))
      .then((doc: DataSourceDocument) => {
        if (
          allowedOrganizations &&
          allowedOrganizations.length > 0 &&
          !allowedOrganizations.includes(doc.toObject().publisherId)
        ) {
          res.status(403).send();

          return;
        }

        res.status(204).send();
        logger.info(`Harvest DataSource with url ${doc.url}`);
        publishDataSource(doc);
      })
      .catch(next);
  }
} as ResourceHandlerMap;
