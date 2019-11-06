import AsyncApiValidator from 'asyncapi-validator';

export const validator = AsyncApiValidator.fromSource(
  'https://raw.githubusercontent.com/Informasjonsforvaltning/fdk-message-broker/master/fdk-message-broker.yaml'
);
