import { RequestHandler } from 'express';
import omit from 'lodash/omit';
import { DataSourceDocument, DataSourceModel } from './data-source.model';
import { NotFoundHttpError } from './lib/http-error';
import uuidv4 from 'uuid/v4';
import { elseThrow } from './lib/else-throw';

interface ResourceHandlerMap {
  create: RequestHandler;
  getById: RequestHandler;
  getAll: RequestHandler;
  update: RequestHandler;
}

export const dataSourceHandlers: ResourceHandlerMap = {
  create: (req, res, next) => {
    const data = req.body;
    data.id = uuidv4(); //todo domain logic should be in domain class, see if it is possible to do it in mongoose

    new DataSourceModel(data)
      .save()
      .then(datasource => {
        res
          .location(datasource.id)
          .status(201)
          .send();
      })
      .catch(next);
  },

  update: (req, res, next) => {
    const { id } = req.params;
    const data = omit(req.body, 'id');
    DataSourceModel.findOneAndUpdate({ id }, data, { new: true })
      .then(elseThrow<DataSourceDocument>(() => new NotFoundHttpError()))
      .then(doc => res.status(200).send(doc.toObject()))
      .catch(next);
  },

  getById: (req, res, next) => {
    const { id } = req.params;
    DataSourceModel.findOne({ id })
      .then(elseThrow<DataSourceDocument>(() => new NotFoundHttpError()))
      .then(doc => res.status(200).send(doc.toObject()))
      .catch(next);
  },

  getAll: (req, res, next) => {
    DataSourceModel.find(req.query)
      .then(docs => res.send(docs.map(doc => doc.toObject())))
      .catch(next);
  }
};
