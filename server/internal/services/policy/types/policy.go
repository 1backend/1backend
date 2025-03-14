/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package policy_svc

import (
	openapi "github.com/openorch/openorch/clients/go"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Entity that is being rate limited/policed
type Entity string

const (
	EntityUserID Entity = "userId"
	EntityIP     Entity = "ip"
)

// Scope is the bucket by which we aggregate
type Scope string

const (
	ScopeEndpoint Scope = "endpoint"
	ScopeGlobal   Scope = "global"
)

type TemplateId string

const (
	TemplateIdRateLimit TemplateId = "rate-limit"
	TemplateIdBlocklist TemplateId = "blocklist"
)

// Parameters for Rate Limit policy instance
type RateLimitParameters struct {
	MaxRequests int    `json:"maxRequests" example:"10"`
	TimeWindow  string `json:"timeWindow"  example:"1m"`
	Entity      Entity `json:"entity"      example:"userId"`
	Scope       Scope  `json:"scope"       example:"endpoint"`
}

type BlocklistParameters struct {
	BlockedIPs []string `json:"blockedIPs"`
}

type Template struct {
	Id          TemplateId `json:"id"          example:"rate-limit"`
	Name        string     `json:"name"        example:"Rate Limit"`
	Description string     `json:"description" example:"Limits the number of requests based on user ID or IP address."`
}

func (t *Template) GetId() string {
	return string(t.Id)
}

var RateLimitPolicyTemplate = Template{
	Id:          TemplateIdRateLimit,
	Name:        "Rate Limit",
	Description: "Limits the number of requests. Aggregate by UserID or IP address.",
}

var BlocklistTemplate = Template{
	Id:          TemplateIdBlocklist,
	Name:        "Blocklist",
	Description: "Block access by IP, UserID and other parameters.",
}

type Instance struct {
	Id         string     `json:"id"`
	Endpoint   string     `json:"endpoint"   example:"/user-svc/register"`
	TemplateId TemplateId `json:"templateId" example:"rate-limit"         binding:"required"`

	Parameters *Parameters `json:"parameters" binding:"required"`
}

type Parameters struct {
	RateLimit *RateLimitParameters `json:"rateLimit,omitempty"`
	Blocklist *BlocklistParameters `json:"blocklist,omitempty"`
}

func (t *Instance) GetId() string {
	return t.Id
}

type UpsertInstanceRequest struct {
	Instance *Instance
}

type UpsertInstanceResponse struct {
}

type CheckRequest struct {
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
	Ip       string `json:"ip"`
	UserId   string `json:"userId"`
}

type CheckResponse struct {
	Allowed bool `json:"allowed" binding:"required"`
}

var (
	PermissionTemplateView = openapi.UserSvcPermission{
		Id:   openapi.PtrString("policy-svc:template:view"),
		Name: openapi.PtrString("Policy Svc - Template View"),
	}

	PermissionTemplateCreate = openapi.UserSvcPermission{
		Id:   openapi.PtrString("policy-svc:template:create"),
		Name: openapi.PtrString("Policy Svc - Template Create"),
	}

	PermissionTemplateEdit = openapi.UserSvcPermission{
		Id:   openapi.PtrString("policy-svc:template:edit"),
		Name: openapi.PtrString("Policy Svc - Template Edit"),
	}

	PermissionTemplateDelete = openapi.UserSvcPermission{
		Id:   openapi.PtrString("policy-svc:template:delete"),
		Name: openapi.PtrString("Policy Svc - Template Delete"),
	}

	PermissionInstanceView = openapi.UserSvcPermission{
		Id:   openapi.PtrString("policy-svc:instance:view"),
		Name: openapi.PtrString("Policy Svc - Instance View"),
	}

	PermissionInstanceCreate = openapi.UserSvcPermission{
		Id:   openapi.PtrString("policy-svc:instance:create"),
		Name: openapi.PtrString("Policy Svc - Instance Create"),
	}

	PermissionInstanceEdit = openapi.UserSvcPermission{
		Id:   openapi.PtrString("policy-svc:instance:edit"),
		Name: openapi.PtrString("Policy Svc - Instance Edit"),
	}

	PermissionInstanceDelete = openapi.UserSvcPermission{
		Id:   openapi.PtrString("policy-svc:instance:delete"),
		Name: openapi.PtrString("Policy Svc - Instance Delete"),
	}

	PermissionCheckView = openapi.UserSvcPermission{
		Id:   openapi.PtrString("policy-svc:check:view"),
		Name: openapi.PtrString("Policy Svc - Check View"),
	}

	AdminPermissions = []openapi.UserSvcPermission{
		PermissionTemplateView,
		PermissionTemplateCreate,
		PermissionTemplateEdit,
		PermissionTemplateDelete,
		PermissionInstanceView,
		PermissionInstanceCreate,
		PermissionInstanceEdit,
		PermissionInstanceDelete,
	}

	UserPermissions = []openapi.UserSvcPermission{
		PermissionCheckView,
	}
)
