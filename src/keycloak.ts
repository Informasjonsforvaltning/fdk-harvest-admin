import Keycloak, { Token } from 'keycloak-connect';
import config from 'config';
import { RequestHandler } from 'express';

const {
  config: keycloakConfig,
  'allowed-authority': allowedAuthority
} = config.get('keycloak');

const keycloakInstance = new Keycloak({}, keycloakConfig);

const checkAuthority = ({ content }: Token): boolean => {
  const { authorities } = content as any;
  return authorities && authorities.split(',').includes(allowedAuthority);
};

export const enforceAuthority = (): RequestHandler => {
  return keycloakInstance.protect(checkAuthority);
};

export default keycloakInstance;
