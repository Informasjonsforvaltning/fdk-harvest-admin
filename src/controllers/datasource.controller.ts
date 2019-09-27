import { Request, Response } from 'express';
import BaseController from './base.controller';

import { DatasourceModel } from '../models/datasource.model';

export class DatasourceController extends BaseController {
  constructor() {
    super();
    this.init();
  }

  private init(): void {
    this.router.get(`/`, this.getAll);
    this.router.post(`/`, this.create);

    this.router.get(`/:id`, this.getById);
    this.router.put(`/:id`, this.update);
  }

  // TODO(chlenix): validate input
  private create(req: Request, res: Response): void {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    new DatasourceModel(req.body).save().then((datasource: any) => {
      res
        .location(datasource.url)
        .status(201)
        .send();
    });
  }

  // TODO(chlenix): validate input
  private update(req: Request, res: Response): void {
    const { id } = req.query;
    DatasourceModel.findByIdAndUpdate(id, req.body).then(() => {
      res.status(200).send();
    });
  }

  // TODO(chlenix): validate input
  private getById(req: Request, res: Response): void {
    const { id } = req.params;
    DatasourceModel.findById(id).then(datasource => {
      res.status(200).send(datasource);
    });
  }

  // TODO(chlenix): validate input
  private getAll(_req: Request, res: Response): void {
    // TODO(chlenix): add optional urlparam filtering

    DatasourceModel.find().then(datasources => {
      res.send(datasources);
    });
  }
}
