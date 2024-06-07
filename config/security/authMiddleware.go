package security

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"

	gocloak "github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func validateTokenAndParseAuthorities(token string) (string, int) {
	client := gocloak.NewClient(env.KeycloakHost())

	ctx := context.Background()
	_, claims, err := client.DecodeAccessToken(ctx, token, "fdk")

	authorities := ""
	errStatus := http.StatusOK

	if err != nil {
		errStatus = http.StatusUnauthorized
	} else if claims == nil {
		errStatus = http.StatusForbidden
	} else {
		var v = jwt.NewValidator(
			jwt.WithLeeway(5*time.Second),
			jwt.WithAudience(env.SecurityValues.TokenAudience),
		)
		validError := v.Validate(claims)
		if validError != nil {
			errStatus = http.StatusForbidden
		}

		authClaim := (*claims)["authorities"]
		if authClaim != nil {
			authorities = authClaim.(string)
		}
	}

	return authorities, errStatus
}

func hasSystemAdminRole(authorities string) bool {
	sysAdminAuth := env.SecurityValues.SysAdminAuth
	return strings.Contains(authorities, sysAdminAuth)
}

func hasOrganizationRole(authorities string, org string, role string) bool {
	orgAdminAuth := fmt.Sprintf("%s:%s:%s", env.SecurityValues.OrgType, org, role)
	return strings.Contains(authorities, orgAdminAuth)
}

func hasAnyOrgAuth(authorities string) bool {
	return strings.Contains(authorities, env.SecurityValues.OrgType)
}

func AuthenticateAndCheckPermissions() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorities, status := validateTokenAndParseAuthorities(c.GetHeader("Authorization"))

		if status != http.StatusOK {
			respondWithError(c, status, http.StatusText(status))
		} else if !(hasSystemAdminRole(authorities) ||
			hasOrganizationRole(authorities, c.Param("org"), env.SecurityValues.AdminPermission) ||
			hasOrganizationRole(authorities, c.Param("org"), env.SecurityValues.WritePermission)) {
			respondWithError(c, http.StatusForbidden, http.StatusText(http.StatusForbidden))
		}

		c.Next()
	}
}

func AuthenticateAnyOrgOrSysAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorities, status := validateTokenAndParseAuthorities(c.GetHeader("Authorization"))

		if status != http.StatusOK {
			respondWithError(c, status, http.StatusText(status))
		} else if !(hasSystemAdminRole(authorities) || hasAnyOrgAuth(authorities)) {
			respondWithError(c, http.StatusForbidden, http.StatusText(http.StatusForbidden))
		}

		c.Next()
	}
}

func AuthenticateApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") != env.ApiKey() {
			respondWithError(c, http.StatusForbidden, http.StatusText(http.StatusForbidden))
		}

		c.Next()
	}
}
