import { Router } from 'express';

export default abstract class BaseController {
  protected router: Router;

  constructor() {
    this.router = Router();
  }

  public get Router(): Router {
    return this.router;
  }
}
