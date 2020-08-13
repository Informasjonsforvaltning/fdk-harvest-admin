import Keycloak, { SpecHandler } from 'keycloak-connect';
import config from 'config';
import { RequestHandler } from 'express';

const { config: keycloakConfig } = config.get('keycloak');

const keycloakInstance = new Keycloak({}, keycloakConfig);

const checkAuthority: SpecHandler = ({ content }, _, res) => {
  const { authorities: authoritiesString } = content as any;

  const authorities: string[] = authoritiesString
    ? authoritiesString.split(',')
    : [];

  const hasSystemAdminRootPermission = authorities.includes(
    'system:root:admin'
  );

  const hasOrganizationAdminPermissions = authorities.some(authority =>
    authority.match(/^organization:\d{9}:admin$/)
  );

  if (hasOrganizationAdminPermissions) {
    res.locals.allowedOrganizations = authorities.reduce(
      (previous, current) => {
        const organizationId = current.match(/^organization:(\d{9}):admin$/);

        return organizationId ? [...previous, organizationId[1]] : previous;
      },
      []
    );
  }

  return hasSystemAdminRootPermission || hasOrganizationAdminPermissions;
};

export const enforceAuthority = (): RequestHandler => {
  return keycloakInstance.protect(checkAuthority);
};

export default keycloakInstance;
