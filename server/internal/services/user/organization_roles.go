package userservice

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/1backend/1backend/sdk/go/auth"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

var ErrOrganizationAdminRequired = errors.New("caller is not an admin of the organization")
var ErrMembershipRoleNotOwned = errors.New("membership role is not owned by caller")
var ErrMembershipRoleOutsideOrganizationScope = errors.New("membership role is outside organization scope")
var ErrMembershipRolesEmpty = errors.New("membership roles cannot be empty")

func orgRolePrefix(organizationId string) string {
	return fmt.Sprintf("user-svc:org:{%s}:", organizationId)
}

func orgUserRole(organizationId string) string {
	return orgRolePrefix(organizationId) + "user"
}

func orgAdminRole(organizationId string) string {
	return orgRolePrefix(organizationId) + "admin"
}

func isOrgScopedRole(role string) bool {
	return strings.HasPrefix(role, "user-svc:org:{")
}

func isOrgScopedRoleForOrganization(organizationId string, role string) bool {
	return strings.HasPrefix(role, orgRolePrefix(organizationId))
}

func normalizeMembershipRoles(organizationId string, roles []string) ([]string, error) {
	if roles == nil {
		roles = []string{orgUserRole(organizationId)}
	}
	if len(roles) == 0 {
		return nil, ErrMembershipRolesEmpty
	}

	roleSet := map[string]struct{}{}
	for _, roleId := range roles {
		if !isOrgScopedRoleForOrganization(organizationId, roleId) {
			return nil, fmt.Errorf("%w: %q", ErrMembershipRoleOutsideOrganizationScope, roleId)
		}
		roleSet[roleId] = struct{}{}
		if roleId == orgAdminRole(organizationId) {
			roleSet[orgUserRole(organizationId)] = struct{}{}
		}
	}

	normalized := make([]string, 0, len(roleSet))
	if _, ok := roleSet[orgUserRole(organizationId)]; ok {
		normalized = append(normalized, orgUserRole(organizationId))
		delete(roleSet, orgUserRole(organizationId))
	}

	remaining := make([]string, 0, len(roleSet))
	for roleId := range roleSet {
		remaining = append(remaining, roleId)
	}
	sort.Strings(remaining)
	normalized = append(normalized, remaining...)

	return normalized, nil
}

func hasOrganizationAdminAccess(Roles []string, organizationId string) bool {
	return contains(Roles, user.RoleAdmin) || contains(Roles, orgAdminRole(organizationId))
}

func validateMembershipRoleOwnership(callerSlug string, callerRoles []string, roles []string) error {
	effectiveClaims := &auth.Claims{
		Slug:  callerSlug,
		Roles: callerRoles,
	}

	for _, roleId := range roles {
		if !auth.OwnsRole(effectiveClaims, roleId) {
			return fmt.Errorf("%w: %q", ErrMembershipRoleNotOwned, roleId)
		}
	}

	return nil
}
