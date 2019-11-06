import AsyncApiValidator from 'asyncapi-validator';

export const validator = AsyncApiValidator.fromSource(
  'https://github.com/Informasjonsforvaltning/fdk-message-broker/blob/master/fdk-message-broker.yaml'
);
