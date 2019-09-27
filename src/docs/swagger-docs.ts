import swaggerUI from 'swagger-ui-express';
import { Express } from 'express';
import YAML from 'yamljs';
import path from 'path';

export default (app: Express): void => {
  // TODO(chlenix): replace yamljs with js-yaml
  const specFile = path.join(__dirname, '../specs', 'fdk-harvest-adm.yaml');
  YAML.load(specFile, specification => {
    app.use('/api-docs', swaggerUI.serve, swaggerUI.setup(specification));
  });
};
