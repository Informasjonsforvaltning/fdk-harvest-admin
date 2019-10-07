import { RequestHandler } from 'express';
import omit from 'lodash/omit';

import { publishToQueue } from './rabbitmq/rabbitmq';
import { DataSourceDocument, DataSourceModel } from './data-source.model';
import { NotFoundHttpError } from './lib/http-error';
import { elseThrow } from './lib/else-throw';

interface ResourceHandlerMap {
  create: RequestHandler;
  getById: RequestHandler;
  getAll: RequestHandler;
  update: RequestHandler;
}

export const createDataSourceHandlers = (): ResourceHandlerMap => {
  return {
    create: (req, res, next): void => {
      const data = omit(req.body, 'id');
      new DataSourceModel(data)
        .save()
        .then(doc => {
          res
            .location(doc.id)
          .status(201)
          .send();

        return doc;
      })
      .then(({ publisherId = '', dataSourceType = '' }) => {
        publishToQueue({
          orgId: publisherId,
          catalogId: publisherId,
          datasourceType: dataSourceType
        });
      })
      .catch(next);
  },

    update: (req, res, next): void => {
      const { id } = req.params;
      const data = omit(req.body, 'id');
      DataSourceModel.findOneAndUpdate({ id }, data, { new: true })
        .then(elseThrow<DataSourceDocument>(() => new NotFoundHttpError()))
        .then(doc => {
        res.status(200).send(doc.toObject());
        return doc;
      })
      .then(({ publisherId = '', dataSourceType = '' }) => {
        publishToQueue({
          orgId: publisherId,
          catalogId: publisherId,
          datasourceType: dataSourceType
        });
      })
        .catch(next);
    },

    getById: (req, res, next): void => {
      const { id } = req.params;
      DataSourceModel.findOne({ id })
        .then(elseThrow<DataSourceDocument>(() => new NotFoundHttpError()))
        .then(doc => res.status(200).send(doc.toObject()))
        .catch(next);
    },

    getAll: (req, res, next): void => {
      DataSourceModel.find(req.query)
        .then(docs => res.send(docs.map(doc => doc.toObject())))
        .catch(next);
    }
  };
};
