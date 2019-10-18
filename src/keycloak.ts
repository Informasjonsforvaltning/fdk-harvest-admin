import Keycloak from 'keycloak-connect';
import config from 'config';

export default new Keycloak({}, config.get('keycloak'));
