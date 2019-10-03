import { ErrorRequestHandler } from 'express';
import { HttpError } from './http-error';

export const commonErrorHandler: ErrorRequestHandler = (
  err,
  _req,
  res,
  _next //eslint-disable-line @typescript-eslint/no-unused-vars
) => {
  const error =
    err instanceof HttpError
      ? err
      : new HttpError(err, err.status || err.statusCode);
  res.status(error.status).json({
    message: err.message,
    errors: err.errors
  });
};
