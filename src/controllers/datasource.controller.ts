/* eslint-disable @typescript-eslint/no-explicit-any */
import { Request, Response, NextFunction } from 'express';
import BaseController from './base.controller';
import { OpenApiValidator } from 'express-openapi-validate';

import { DatasourceModel } from '../models/datasource.model';
import { Utils } from '../utils/utils';

export class DatasourceController extends BaseController {
  // exclude mongodb's _id field
  static DTO = 'id url publisher description dataSourceType -_id';
  static PATH = '/datasources';

  private validator: OpenApiValidator;

  constructor() {
    super();
    this.init();
  }

  private init(): void {
    this.validator = new OpenApiValidator(Utils.readOpenApi());

    this.router.get(
      `/`,
      this.validator.validate('get', DatasourceController.PATH),
      this.getAll
    );

    this.router.post(
      `/`,
      this.validator.validate('post', DatasourceController.PATH),
      this.create
    );

    this.router.get(
      `/:id`,
      this.validator.validate('get', `${DatasourceController.PATH}/{id}`),
      this.getById
    );

    // TODO: replace with Patch
    this.router.put(
      `/:id`,
      this.validator.validate('put', `${DatasourceController.PATH}/{id}`),
      this.update
    );

    this.router.use(this.errorHandler);
  }

  private create(req: Request, res: Response, next: NextFunction): void {
    new DatasourceModel(req.body)
      .save()
      .then((datasource: any) => {
        res
          .location(datasource.url)
          .status(201)
          .send();
      })
      .catch(next);
  }

  private update(req: Request, res: Response, next: NextFunction): void {
    // id from body used for update data
    const { id, dataSourceType, url, publisher, description } = req.body;

    const filter = {
      id: req.params.id
    };

    const update = {
      id: id,
      dataSourceType: dataSourceType,
      url: url,
      publisher: publisher,
      description: description
    };

    // weird possible case: Pathvariable 'id' and 'id' from body may differ,
    // and it would still be a valid request per spec file
    DatasourceModel.findOneAndUpdate(filter, update)
      .then(() => {
        res.status(200).send();
      })
      .catch(next);
  }

  // TODO(chlenix): validate input
  private getById(req: Request, res: Response, next: NextFunction): void {
    const { id } = req.params;

    DatasourceModel.findOne({ id: id }, DatasourceController.DTO)
      .then(datasource => {
        const code = datasource ? 200 : 404;
        res.status(code).send(datasource);
      })
      .catch(next);
  }

  private getAll(req: Request, res: Response, next: NextFunction): void {
    const { publisher, dataSourceType } = req.query;

    const filter = {
      publisher: publisher || /.*/,
      dataSourceType: dataSourceType || /.*/
    };

    DatasourceModel.find(filter, DatasourceController.DTO)
      .then(datasources => {
        const code = datasources.length ? 200 : 404;
        res.status(code).send(datasources);
      })
      .catch(next);
  }

  private static formatInternal(err: Error): any {
    const { name, message, stack } = err;
    return {
      name: name,
      message: message || 'unexpected error',
      stack: stack,
      statusCode: 500
    };
  }

  private errorHandler(
    err: Error,
    _req: Request,
    _res: Response,
    next: NextFunction
  ): void {
    // simply forward ValidationError
    // otherwise create and send a new error

    next(
      err.name === 'ValidationError'
        ? err
        : DatasourceController.formatInternal(err)
    );
  }
}
