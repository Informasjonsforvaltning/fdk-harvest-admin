export class HttpError extends Error {
  status: number;

  constructor(message = 'Unknown error', status = 500) {
    super(message);
    this.status = status;
  }
}

export class NotFoundHttpError extends HttpError {
  constructor(message = 'Not found') {
    super(message, 404);
  }
}
