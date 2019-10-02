import express from 'express';

export const commonErrorHandler = (
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  err: any,
  _req: express.Request,
  res: express.Response,
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  _next: express.NextFunction
): void => {
  // format error
  res.status(err.status).json({
    message: err.message,
    errors: err.errors
  });
};
